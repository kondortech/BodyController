package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/kirvader/BodyController/services/food/ingredient/proto"
)

type server struct {
	pb.UnimplementedInventoryServer
}

func (s *server) GetBookList(_ context.Context, _ *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	return &pb.GetBookResponse{
		Book: &pb.Book{
			Title:     "title",
			Author:    "author",
			PageCount: 322,
			Language:  "eng",
		},
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterInventoryServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
