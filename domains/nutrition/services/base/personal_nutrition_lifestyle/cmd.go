package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/kirvader/BodyController/domains/nutrition/services/base/personal_nutrition_lifestyle/src"
	"github.com/kirvader/BodyController/pkg/utils"
)

func main() {
	servicePort := utils.GetEnvWithDefault("SERVICE_PORT", "10000")
	serviceURI := fmt.Sprintf(":%s", servicePort)
	log.Println("service uri: ", serviceURI)

	listener, err := net.Listen("tcp", serviceURI)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	svc, close, err := src.NewPersonalNutritionLifestyleService(ctx)
	if err != nil {
		panic(err)
	}
	defer close()

	if err := svc.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
