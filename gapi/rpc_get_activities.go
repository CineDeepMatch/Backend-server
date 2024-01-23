package gapi

import (
	"context"

	"github.com/CineDeepMatch/Backend-server/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetActivities(ctx context.Context, req *pb.GetActivitiesRequest) (*pb.GetActivitiesResponse, error) {
	userId := uuid.MustParse(req.GetUserId())

	activities, err := server.store.GetActivitiesByUserId(ctx, userId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user's activities: %s", err)
	}

	rsp := &pb.GetActivitiesResponse{
		Activities: convertActivities(activities),
	}

	return rsp, nil
}
