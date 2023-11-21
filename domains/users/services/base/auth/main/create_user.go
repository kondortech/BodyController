package main

import (
	"context"
	"log"

	pbAuth "github.com/kirvader/BodyController/domains/users/services/base/auth/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (svc *AuthService) CreateUser(ctx context.Context, req *pbAuth.CreateUserRequest) (*pbAuth.CreateUserResponse, error) {
	userCredentialsCollection := svc.mongoClient.Database("BodyController").Collection("UserCredentials")
	indexName, err := userCredentialsCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "Username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "error occured when index was created")
	}
	log.Printf("unique index created: %s", indexName)

	// var creds pbAuth.UserCredentials
	// userCredentialsCollection.FindOne(ctx, bson.D{{Key: "Username", Value: "kk"}}).Decode(&creds)

	// log.Print(creds)

	// todo check if request is valid

	userCredentialsCollection.InsertOne(ctx, req.GetUserCredentials())

	return &pbAuth.CreateUserResponse{}, nil
}
