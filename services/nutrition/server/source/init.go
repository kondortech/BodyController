package source

import (
	"context"
	"errors"
	"fmt"
	"github.com/kirvader/BodyController/pkg/utils"
	pb "github.com/kirvader/BodyController/services/nutrition/proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var Cmd = &cobra.Command{
	Use:   "nutrition",
	Short: "Runs nutrition server with variables from env",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serveFromEnv(cmd.Context())
	},
}

func serveFromEnv(ctx context.Context) error {
	servicePort, err := utils.LookupEnv("SERVICE_PORT")
	if err != nil {
		return err
	}

	mongodbHost, err := utils.LookupEnv("MONGODB_HOST")
	if err != nil {
		return err
	}
	mongodbPort, err := utils.LookupEnv("MONGODB_PORT")
	if err != nil {
		return err
	}

	mongodbUsername, err := utils.LookupEnv("MONGODB_USERNAME")
	if err != nil {
		return err
	}
	mongodbPassword, err := utils.LookupEnv("MONGODB_PASSWORD")
	if err != nil {
		return err
	}

	rabbitMQHost, err := utils.LookupEnv("RABBITMQ_HOST")
	if err != nil {
		return err
	}
	rabbitMQPort, err := utils.LookupEnv("RABBITMQ_PORT")
	if err != nil {
		return err
	}

	rabbitMQUser, err := utils.LookupEnv("RABBITMQ_USERNAME")
	if err != nil {
		return err
	}
	rabbitMQPassword, err := utils.LookupEnv("RABBITMQ_PASSWORD")
	if err != nil {
		return err
	}

	return serve(ctx, servicePort, mongodbHost, mongodbPort, mongodbUsername, mongodbPassword, rabbitMQHost, rabbitMQPort, rabbitMQUser, rabbitMQPassword)
}

type service struct {
	mongoClient  *mongo.Client
	rabbitMQConn *amqp.Connection

	pb.UnimplementedNutritionServer
}

// TODO refactor used components
func serve(ctx context.Context,
	servicePort,
	mongodbHost, mongodbPort, mongodbUsername, mongodbPassword,
	rabbitmqHost, rabbitmqPort, rabbitmqUsername, rabbitmqPassword string) (err error) {

	opts := mongoOptions.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%s/", mongodbHost, mongodbPort)).
		// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
		SetServerAPIOptions(mongoOptions.ServerAPI(mongoOptions.ServerAPIVersion1))
	// TODO add auth

	// Create a new client and connect to the server
	mongoClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		return fmt.Errorf("failed to connect to mongodb: %w", err)
	}
	defer func() {
		if disconnectErr := mongoClient.Disconnect(ctx); disconnectErr != nil {
			err = errors.Join(err, fmt.Errorf("failed to close mongodb: %w", disconnectErr))
		}
	}()

	// Send a ping to confirm a successful connection
	if pingErr := mongoClient.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); pingErr != nil {
		return fmt.Errorf("failed to ping mongodb: %w", pingErr)
	}

	rabbitmqConn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", rabbitmqUsername, rabbitmqPassword, rabbitmqHost, rabbitmqPort))
	if err != nil {
		return fmt.Errorf("failed to connect to rabbitmq: %w", err)
	}
	defer func() {
		closeErr := rabbitmqConn.Close()
		if closeErr != nil {
			err = errors.Join(err, fmt.Errorf("failed to close rabbitmq connection: %w", closeErr))
		}
	}()

	fmt.Println("local addr mq: ", rabbitmqConn.LocalAddr())
	fmt.Println("remote addr mq: ", rabbitmqConn.RemoteAddr())

	svc := &service{
		mongoClient:  mongoClient,
		rabbitMQConn: rabbitmqConn,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNutritionServer(grpcServer, svc)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", servicePort))
	if err != nil {
		return fmt.Errorf("failed to start listening on %q: %w", fmt.Sprintf(":%s", servicePort), err)
	}
	return grpcServer.Serve(listener)
}
