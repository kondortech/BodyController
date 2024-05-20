package src

import (
	"context"
	"fmt"
	"net"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/kirvader/BodyController/internal/db"
	"github.com/kirvader/BodyController/pkg/utils"
	pb "github.com/kirvader/BodyController/services/nutrition/proto"
)

type Service struct {
	mongoClient     *mongo.Client
	workerChannelMQ *amqp.Channel

	pb.UnimplementedNutritionServer
}

func Serve(ctx context.Context) error {
	mongoClient, closeMongoClient, err := db.InitMongoDBClientFromENV(ctx)
	if err != nil {
		return err
	}
	defer closeMongoClient()

	if err = db.PingMongoDb(ctx, mongoClient); err != nil {
		return err
	}

	servicePort := utils.GetEnvWithDefault("SERVICE_PORT", "50051")
	serviceURI := fmt.Sprintf(":%s", servicePort)
	listener, err := net.Listen("tcp", serviceURI)
	if err != nil {
		panic(err)
	}

	mqConn, err := amqp.Dial("amqp://guest:guest@nutrition-message-broker-rabbitmq:5672/")
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %s", err)
	}
	defer mqConn.Close()

	mqChannel, err := mqConn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel on amqp connection: %s", err)
	}
	defer mqChannel.Close()

	svc := &Service{
		mongoClient:     mongoClient,
		workerChannelMQ: mqChannel,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNutritionServer(grpcServer, svc)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener)
}
