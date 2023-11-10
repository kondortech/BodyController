package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/kirvader/BodyController/services/nutrition/ingredient/proto"
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

	// TODO make it depend on the way of using it
	reflection.Register(s)

	pb.RegisterInventoryServer(s, &server{})
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
