package gapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/CineDeepMatch/Backend-server/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthInterceptor struct {
	tokenMaker token.Maker
}

func NewAuthInterceptor(tokenMaker token.Maker) *AuthInterceptor {
	return &AuthInterceptor{
		tokenMaker,
	}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing metadata")
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return fmt.Errorf("missing authorization header")
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return fmt.Errorf("invalid authorization header format")
	}
	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return fmt.Errorf("unsupport authorization type: %s ", authType)
	}

	accessToken := fields[1]
	_, err := interceptor.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return fmt.Errorf("invalid access token: %s", err)
	}
	return nil

}
