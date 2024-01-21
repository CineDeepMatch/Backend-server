package gapi

import (
	"context"
	"strings"

	db "github.com/CineDeepMatch/Backend-server/db/sqlc"
	"github.com/CineDeepMatch/Backend-server/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateFavMovies(ctx context.Context, req *pb.CreateFavMoviesRequest) (*pb.CreateFavMoviesResponse, error) {
	userId := uuid.MustParse(req.GetUserId())

	arg := db.CreateFavMoviesParams{
		UserID: userId,
		Movies: strings.Join(req.GetMovieIds(), " "),
	}

	favMovies, err := server.store.CreateFavMovies(ctx, arg)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create activity: %s", err)
	}

	rsp := &pb.CreateFavMoviesResponse{
		FavMovies: convertFavMovies(favMovies),
	}
	return rsp, nil
}
