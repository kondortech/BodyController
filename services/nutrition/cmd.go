package main

import (
	"context"
	"log"

	server "github.com/kirvader/BodyController/services/nutrition/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := server.Serve(ctx); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
