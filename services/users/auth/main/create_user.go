package main

import (
	"context"

	pbAuth "github.com/kirvader/BodyController/services/auth/proto"
)

// CreateUser implements proto.AuthServer.
func (svc *AuthService) CreateUser(context.Context, *pbAuth.CreateUserRequest) (*pbAuth.CreateUserResponse, error) {
	panic("unimplemented")
}
