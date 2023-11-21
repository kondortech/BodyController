package main

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	pbAuth "github.com/kirvader/BodyController/domains/users/services/base/auth/proto"
	user "github.com/kirvader/BodyController/models/users"
)

// this application is not in production yet so this is just a dummy secret
// TODO hide it
const secretKey = "abracadabra this is secret key"

func (svc *AuthService) LogIn(ctx context.Context, req *pbAuth.LogInRequest) (*pbAuth.LogInResponse, error) {
	// TODO validate user

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username":        req.UserCredentials.Username,
			"expiration_date": time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	authToken, err := token.SignedString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("generating of authentication token failed: %w", err)
	}
	return &pbAuth.LogInResponse{
		LoggedUserInfo: &user.LoggedUserInfo{
			Username: req.UserCredentials.Username,
			Token:    authToken,
		},
	}, nil
}
