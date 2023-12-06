package src

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	pb "github.com/kirvader/BodyController/domains/security/services/aggregation/proto"
	pbAuth "github.com/kirvader/BodyController/domains/security/services/base/auth/proto"
	"github.com/kirvader/BodyController/pkg/utils"
)

type SecurityService struct {
	authClient pbAuth.AuthClient

	pb.UnimplementedSecurityServer
}

func NewSecurityService(ctx context.Context) (*SecurityService, func(), error) {
	authServiceClientPort := utils.GetEnvWithDefault("BASE_AUTH_IP", "0.0.0.0")
	authServiceClientIP := utils.GetEnvWithDefault("BASE_AUTH_PORT", "10001")

	authServiceURI := fmt.Sprintf("%s:%s", authServiceClientPort, authServiceClientIP)
	log.Printf("base-auth uri: %s", authServiceURI)

	conn, err := grpc.Dial(authServiceURI, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &SecurityService{
			authClient: pbAuth.NewAuthClient(conn),
		}, func() {
			conn.Close()
		}, nil
}

func (svc *SecurityService) Serve(listener net.Listener) error {
	grpcServer := grpc.NewServer()
	pb.RegisterSecurityServer(grpcServer, svc)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener)
}
