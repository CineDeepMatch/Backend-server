package gapi

import (
	"context"

	db "github.com/CineDeepMatch/Backend-server/db/sqlc"
	"github.com/CineDeepMatch/Backend-server/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateActivity(ctx context.Context, req *pb.CreateActivityRequest) (*pb.CreateActivityResponse, error) {
	userId := uuid.MustParse(req.GetUserId())

	// authPayload, err := server.authorizeUser(ctx)

	// if err != nil {
	// 	return nil, unauthenticatedError(err)
	// }

	// if authPayload.UserId != userId {
	// 	return nil, status.Errorf(codes.PermissionDenied, "cannot create other user's activity")
	// }

	activityId, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	arg := db.CreateActivityParams{
		ID:            activityId,
		UserID:        userId,
		ViewPage:      req.GetViewPage(),
		Duration:      req.GetDuration(),
		PageVisitedAt: req.GetPageVisitedAt().AsTime(),
	}

	activity, err := server.store.CreateActivity(ctx, arg)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create activity: %s", err)
	}

	rsp := &pb.CreateActivityResponse{
		Activity: convertActivity(activity),
	}
	return rsp, nil
}

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