package gapi

import (
	"context"
	"database/sql"
	"strings"

	db "github.com/CineDeepMatch/Backend-server/db/sqlc"
	"github.com/CineDeepMatch/Backend-server/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateFavMovies(ctx context.Context, req *pb.UpdateFavMoviesRequest) (*pb.UpdateFavMoviesResponse, error) {
	userId := uuid.MustParse(req.GetUserId())

	arg := db.UpdateFavMoviesByUserIdParams{
		UserID: userId,
		Movies: strings.Join(req.GetMovieIds(), " "),
	}

	fav_movies, err := server.store.UpdateFavMoviesByUserId(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found: %s", err)

		}
		return nil, status.Errorf(codes.Internal, "failed to update fav movies: %s", err)
	}

	rsp := &pb.UpdateFavMoviesResponse{
		FavMovies: convertFavMovies(fav_movies),
	}
	return rsp, nil
}
