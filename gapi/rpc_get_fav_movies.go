package gapi

import (
	"context"

	"github.com/CineDeepMatch/Backend-server/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetFavMovies(ctx context.Context, req *pb.GetFavMoviesRequest) (*pb.GetFavMoviesResponse, error) {
	userId := uuid.MustParse(req.GetUserId())

	favMovies, err := server.store.GetFavMoviesByUserId(ctx, userId)
	
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user's favorite movies: %s", err)
	}

	rsp := &pb.GetFavMoviesResponse{
		FavMovies: convertFavMovies(favMovies),
	}

	return rsp, nil
}
