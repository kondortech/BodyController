package main

import (
	"context"
	"fmt"

	pbAuth "github.com/kirvader/BodyController/domains/users/services/base/auth/proto"
)

func (svc *AuthService) DeleteUser(ctx context.Context, req *pbAuth.DeleteUserRequest) (*pbAuth.DeleteUserResponse, error) {
	// TODO validate user is authorized

	userCredentialsCollection := svc.mongoClient.Database("BodyController").Collection("UserCredentials")

	if _, err := userCredentialsCollection.DeleteOne(ctx, req.GetUsername()); err != nil {
		return nil, fmt.Errorf("delete user request returned an error: %w", err)
	}

	return &pbAuth.DeleteUserResponse{}, nil
}
