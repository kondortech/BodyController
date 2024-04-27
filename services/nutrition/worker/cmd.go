package main

import (
	"context"
	"log"

	"github.com/kirvader/BodyController/services/nutrition/worker/src"
)

func main() {
	if err := src.InitConsumer(context.Background()); err != nil {
		log.Fatalf("worker crashed with error: %v", err)
	}
	log.Print("worker execution finished")
}
