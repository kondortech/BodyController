package src

import (
	"context"
	"fmt"
	"log"

	"github.com/kirvader/BodyController/internal/db"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/sync/errgroup"
)

func InitConsumer(ctx context.Context) error {
	mongoClient, closeMongoClient, err := db.InitMongoDBClientFromENV(ctx)
	if err != nil {
		return err
	}
	defer closeMongoClient()

	conn, err := amqp.Dial("amqp://guest:guest@nutrition-message-broker-rabbitmq:5672/")
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ingredientMQChannel, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %s", err)
	}
	defer ingredientMQChannel.Close()

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

	recipeMQChannel, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %s", err)
	}
	defer recipeMQChannel.Close()

	recipeConsumer, err := recipeMQChannel.ConsumeWithContext(
		ctx,
		"recipe",
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return fmt.Errorf("failed to create consumer for recipe: %s", err)
	}
	fmt.Println("worker is initialized and consumes recipe messages")

	mealMQChannel, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %s", err)
	}
	defer mealMQChannel.Close()

	mealConsumer, err := mealMQChannel.ConsumeWithContext(
		ctx,
		"meal",
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return fmt.Errorf("failed to create consumer for meal: %s", err)
	}
	fmt.Println("worker is initialized and consumes meal messages")

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
	eg.Go(func() error {
		// TODO make this parallel with pool
		for {
			select {
			case item := <-recipeConsumer:
				err := ProcessRecipeOperation(ctx, mongoClient, item)
				if err != nil {
					log.Printf("recipe operation processing failed: %v", err)
					return err
				}
			case <-egCtx.Done():
				log.Print("context canceled")
				return nil
			}
		}
	})
	eg.Go(func() error {
		// TODO make this parallel with pool
		for {
			select {
			case item := <-mealConsumer:
				err := ProcessMealOperation(ctx, mongoClient, item)
				if err != nil {
					log.Printf("meal operation processing failed: %v", err)
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
