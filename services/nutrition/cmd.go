package main

import (
	"context"
	"log"

	"github.com/kirvader/BodyController/services/nutrition/src"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := src.Serve(ctx); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
