package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// this application is not yet in production so this is just a dummy secret
// TODO hide it
const secretKey = "abracadabra this is secret key"

func GenerateAuthJWT(username string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username":        username,
			"expiration_date": time.Now().Add(duration).Unix(),
		},
	)

	authToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("generating of authentication token failed: %w", err)
	}

	return authToken, nil
}

func CheckAuthJWTToken(jwtToken, username string) error {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return status.Error(codes.InvalidArgument, "Invalid JWT token")
	}
	if !token.Valid {
		return status.Error(codes.InvalidArgument, "Invalid JWT token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return status.Error(codes.InvalidArgument, "Invalid JWT token claims")
	}
	tokenUsername, ok := claims["username"]
	if !ok {
		return status.Error(codes.InvalidArgument, "Missing JWT token claim: username")
	}
	tokenExpirationDate, ok := claims["expiration_date"]
	if !ok {
		return status.Error(codes.InvalidArgument, "Missing JWT token claim: username")
	}

	if tokenUsername != username {
		return status.Error(codes.Unauthenticated, "Auth token has different username")
	}

	if time.Now().Before(tokenExpirationDate.(time.Time)) {
		return status.Error(codes.Unauthenticated, "Auth token has expired")
	}

	return nil
}
