package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/kirvader/BodyController/domains/nutrition/services/aggregation/src"
	"github.com/kirvader/BodyController/pkg/utils"
)

func main() {
	servicePort := utils.GetEnvWithDefault("SERVICE_PORT", "20000")
	serviceURI := fmt.Sprintf(":%s", servicePort)

	listener, err := net.Listen("tcp", serviceURI)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	svc, close, err := src.NewNutritionService(ctx)
	if err != nil {
		panic(err)
	}
	defer close()

	if err := svc.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
