package source

import (
	"context"
	"errors"
	"fmt"
	"github.com/kirvader/BodyController/pkg/utils"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"
)

var Cmd = &cobra.Command{
	Use:   "nutrition-worker",
	Short: "Runs nutrition server worker",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serveFromEnv(cmd.Context())
	},
}

func serveFromEnv(ctx context.Context) error {
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

	return serve(ctx, mongodbHost, mongodbPort, mongodbUsername, mongodbPassword, rabbitMQHost, rabbitMQPort, rabbitMQUser, rabbitMQPassword)
}

type worker struct {
	mongoClient  *mongo.Client
	rabbitMQConn *amqp.Connection
}

func serve(ctx context.Context,
	mongodbHost, mongodbPort, mongodbUsername, mongodbPassword,
	rabbitmqHost, rabbitmqPort, rabbitmqUsername, rabbitmqPassword string) (err error) {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", mongodbUsername, mongodbPassword, mongodbHost, mongodbPort)
	opts := options.Client().ApplyURI(mongoUri).SetServerAPIOptions(serverAPI)
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

	workerService := &worker{
		mongoClient:  mongoClient,
		rabbitMQConn: rabbitmqConn,
	}

	ingredientMQChannel, err := workerService.rabbitMQConn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %s", err)
	}
	defer func() {
		closeErr := ingredientMQChannel.Close()
		if closeErr != nil {
			err = errors.Join(err, fmt.Errorf("failed to close rabbitmq channel: %w", closeErr))
		}
	}()

	ingredientConsumer, err := ingredientMQChannel.ConsumeWithContext(
		ctx,
		"ingredient",
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return fmt.Errorf("failed to create consumer for ingredient: %s", err)
	}
	fmt.Println("worker is initialized and consumes ingredient messages")

	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		// TODO make this parallel with pool
		for {
			select {
			case item := <-ingredientConsumer:
				err := ProcessIngredientOperation(ctx, mongoClient, item)
				if err != nil {
					log.Printf("ingredient operation processing failed: %v", err)
					return err
				}
			case <-egCtx.Done():
				log.Print("context canceled")
				return nil
			}
		}
	})
	return eg.Wait()
}
