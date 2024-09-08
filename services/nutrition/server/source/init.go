package source

import (
	"context"
	"errors"
	"fmt"
	"github.com/kirvader/BodyController/pkg/utils"
	pb "github.com/kirvader/BodyController/services/nutrition/proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

	mongoClientIP, err := utils.LookupEnv("MONGO_IP")
	if err != nil {
		return err
	}
	mongoClientPort, err := utils.LookupEnv("MONGO_PORT")
	if err != nil {
		return err
	}

	rabbitMQUser, err := utils.LookupEnv("RABBITMQ_USER")
	if err != nil {
		return err
	}
	rabbitMQPassword, err := utils.LookupEnv("RABBITMQ_PASSWORD")
	if err != nil {
		return err
	}

	rabbitMQIP, err := utils.LookupEnv("RABBITMQ_IP")
	if err != nil {
		return err
	}
	rabbitMQPort, err := utils.LookupEnv("RABBITMQ_PORT")
	if err != nil {
		return err
	}

	return serve(ctx, servicePort, mongoClientIP, mongoClientPort, rabbitMQUser, rabbitMQPassword, rabbitMQIP, rabbitMQPort)
}

type service struct {
	mongoClient  *mongo.Client
	rabbitMQConn *amqp.Channel

	pb.UnimplementedNutritionServer
}

func serve(ctx context.Context, servicePort, mongoClientIP, mongoClientPort, rabbitMQUser, rabbitMQPassword, rabbitMQIP, rabbitMQPort string) (err error) {
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", mongoClientIP, mongoClientPort)))
	if err != nil {
		return err
	}
	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	serviceURI := fmt.Sprintf(":%s", servicePort)
	listener, err := net.Listen("tcp", serviceURI)
	if err != nil {
		return err
	}

	mqConn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s", rabbitMQUser, rabbitMQPassword, net.JoinHostPort(rabbitMQIP, rabbitMQPort)))
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		closeErr := mqConn.Close()
		if closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}()

	mqChannel, err := mqConn.Channel()
	if err != nil {
		return err
	}
	defer func() {
		closeErr := mqChannel.Close()
		if closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}()

	svc := &service{
		mongoClient:  mongoClient,
		rabbitMQConn: mqChannel,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNutritionServer(grpcServer, svc)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener)
}
