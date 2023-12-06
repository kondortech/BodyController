package src

import (
	"context"
	"fmt"

	pbAuth "github.com/kirvader/BodyController/domains/security/services/base/auth/proto"
)

func (svc *AuthService) CreateUser(ctx context.Context, req *pbAuth.CreateUserRequest) (*pbAuth.CreateUserResponse, error) {
	userCredentialsCollection := svc.mongoClient.Database("BodyController").Collection("UserCredentials")

	mongoUserCredentials, err := req.UserCredentials.ConvertToMongoReadable()
	if err != nil {
		return nil, err
	}

	if _, err := userCredentialsCollection.InsertOne(ctx, mongoUserCredentials); err != nil {
		return nil, fmt.Errorf("insert error: %w", err)
	}

	return &pbAuth.CreateUserResponse{}, nil
}
