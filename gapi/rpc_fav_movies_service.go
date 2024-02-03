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

func (server *Server) GetFavMovies(ctx context.Context, req *pb.GetFavMoviesRequest) (*pb.GetFavMoviesResponse, error) {
	userId := uuid.MustParse(req.GetUserId())

	favMovies, err := server.store.GetFavMoviesByUserId(ctx, userId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user's favorite movies: %s", err)
	}

	favMovieIds := strings.Split(favMovies.Movies, " ")

	movies, err := server.mongoDBStore.GetMovieByIds(ctx, favMovieIds)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get movies: %s", err)
	}

	rsp := &pb.GetFavMoviesResponse{
		FavMovies: convertMovies(movies),
	}

	return rsp, nil
}

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
