package main

import (
	"context"

	"github.com/kirvader/BodyController/services/gateway/server/src"
	"google.golang.org/grpc/grpclog"
)

func main() {
	ctx := context.Background()
	opts := src.Options{
		Addr: ":8080",
		GRPCServer: src.Endpoint{
			Network: "tcp",
			Addr:    "nutrition-service:50001",
		},
		OpenAPIDir: "/generated/openapiv2/services/nutrition/proto/api.swagger.json",
	}
	if err := src.Run(ctx, opts); err != nil {
		grpclog.Fatal(err)
	}
}
