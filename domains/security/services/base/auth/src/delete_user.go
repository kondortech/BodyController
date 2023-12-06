package src

import (
	"context"
	"fmt"

	pbAuth "github.com/kirvader/BodyController/domains/security/services/base/auth/proto"
	"github.com/kirvader/BodyController/internal/auth"
	"go.mongodb.org/mongo-driver/bson"
)

func (svc *AuthService) DeleteUser(ctx context.Context, req *pbAuth.DeleteUserRequest) (*pbAuth.DeleteUserResponse, error) {
	if err := auth.CheckUserIsAuthorized(ctx, req.Username); err != nil {
		return nil, err
	}

	userCredentialsCollection := svc.mongoClient.Database("BodyController").Collection("UserCredentials")

	if _, err := userCredentialsCollection.DeleteOne(ctx, bson.D{{Key: "username", Value: req.Username}}); err != nil {
		return nil, fmt.Errorf("delete user request returned an error: %w", err)
	}

	return &pbAuth.DeleteUserResponse{}, nil
}
