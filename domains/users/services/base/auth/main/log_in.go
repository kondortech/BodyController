package main

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	pbAuth "github.com/kirvader/BodyController/domains/users/services/base/auth/proto"
)

// yet this application is not deployed so secrets don't have any place to be stored
// TODO hide it
const secretKey = "abracadabra this is secret key"

func (svc *AuthService) LogIn(ctx context.Context, req *pbAuth.LogInRequest) (*pbAuth.LogInResponse, error) {
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
		LoggedUserInfo: &pbAuth.LoggedUserInfo{
			Username: req.UserCredentials.Username,
			Token:    authToken,
		},
	}, nil
}
