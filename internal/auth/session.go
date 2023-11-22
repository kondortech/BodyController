package auth

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const AuthHeaderName = "JWTAuthToken"

func WithJWTAuth(ctx context.Context, token string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, AuthHeaderName, token)
}

func getJWTAuthToken(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	tokensList, ok := md[AuthHeaderName]
	if !ok {
		return ""
	}
	if len(tokensList) != 1 {
		return ""
	}
	return tokensList[0]
}

func CheckUserIsAuthorized(ctx context.Context, username string) error {
	return CheckAuthJWTToken(getJWTAuthToken(ctx), username)
}
