package gapi

import (
	"context"
	"errors"
	"time"

	db "github.com/CineDeepMatch/Backend-server/db/sqlc"
	"github.com/CineDeepMatch/Backend-server/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) RenewAccessToken(ctx context.Context, req *pb.RenewAccessTokenRequest) (*pb.RenewAccessTokenResponse, error) {

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized refresh token: %s", err)
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)

	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "session not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "unauthorized refresh token: %s", err)
	}

	if session.IsBlocked {
		return nil, status.Errorf(codes.Unauthenticated, "block session")
	}

	if session.UserID != refreshPayload.UserId {
		return nil, status.Errorf(codes.Unauthenticated, "incorrect session user")
	}

	if session.RefreshToken != req.RefreshToken {
		return nil, status.Errorf(codes.Unauthenticated, "mismatched session token")
	}

	if time.Now().After(session.ExpiresAt) {
		return nil, status.Errorf(codes.Unauthenticated, "expired session")
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		refreshPayload.UserId,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create access token: %s", err)
	}

	rsp := &pb.RenewAccessTokenResponse{
		AccessToken: convertToken(accessToken, accessPayload.ExpiredAt),
	}

	return rsp, nil
}
