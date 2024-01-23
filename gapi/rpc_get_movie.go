package gapi

import (
	"context"

	"github.com/CineDeepMatch/Backend-server/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetMovie(ctx context.Context, req *pb.GetMovieRequest) (*pb.GetMovieResponse, error) {

	movie, err := server.mongoDBStore.GetMovieById(ctx, req.GetMovieId())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get movie: %s", err)
	}

	rsp := &pb.GetMovieResponse{
		Movie: convertMovie(movie),
	}

	return rsp, nil
}
