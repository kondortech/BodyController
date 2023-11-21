package main

import (
	"context"
	"fmt"

	pbAuth "github.com/kirvader/BodyController/domains/users/services/base/auth/proto"
)

func (svc *AuthService) CreateUser(ctx context.Context, req *pbAuth.CreateUserRequest) (*pbAuth.CreateUserResponse, error) {
	userCredentialsCollection := svc.mongoClient.Database("BodyController").Collection("UserCredentials")

	if _, err := userCredentialsCollection.InsertOne(ctx, req.GetUserCredentials()); err != nil {
		return nil, fmt.Errorf("insert error: %w", err)
	}

	return &pbAuth.CreateUserResponse{}, nil
}
