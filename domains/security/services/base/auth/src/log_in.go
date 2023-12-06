package src

import (
	"context"
	"fmt"
	"time"

	"github.com/kirvader/BodyController/domains/security/models"
	pbAuth "github.com/kirvader/BodyController/domains/security/services/base/auth/proto"

	"github.com/kirvader/BodyController/internal/auth"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (svc *AuthService) LogIn(ctx context.Context, req *pbAuth.LogInRequest) (*pbAuth.LogInResponse, error) {
	if err := checkUserCredentials(svc.mongoClient, req.UserCredentials); err != nil {
		return nil, err
	}

	authToken, err := auth.GenerateAuthJWT(req.UserCredentials.Username, time.Hour*24)
	if err != nil {
		return nil, fmt.Errorf("generating of authentication token failed: %w", err)
	}

	return &pbAuth.LogInResponse{
		Token: authToken,
	}, nil
}

func checkUserCredentials(mongoClient *mongo.Client, providedUserCredentials *models.UserCredentials) error {
	userCredentialsCollection := mongoClient.Database("BodyController").Collection("UserCredentials")

	var userCredentialsRecord models.UserCredentials
	err := userCredentialsCollection.FindOne(
		context.TODO(),
		bson.D{{Key: "username", Value: providedUserCredentials.Username}},
		options.FindOne().SetProjection(bson.D{{Key: "username", Value: 1}, {Key: "password", Value: 1}})).
		Decode(&userCredentialsRecord)

	if err != nil {
		return status.Error(codes.InvalidArgument, fmt.Sprintf("provided username doesn't exist: %s", providedUserCredentials.Username))
	}

	if userCredentialsRecord.Password != providedUserCredentials.Password {
		return status.Error(codes.InvalidArgument, "wrong password for provided username")
	}
	return nil
}
