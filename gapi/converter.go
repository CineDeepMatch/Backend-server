package gapi

import (
	"strings"
	"time"

	db "github.com/CineDeepMatch/Backend-server/db/sqlc"
	mongodb "github.com/CineDeepMatch/Backend-server/mongodb/repositories"
	"github.com/CineDeepMatch/Backend-server/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Id:                user.ID.String(),
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreateAt),
	}
}

func convertActivity(activity db.Activity) *pb.Activity {
	return &pb.Activity{
		Id:            activity.ID.String(),
		UserId:        activity.UserID.String(),
		ViewPage:      activity.ViewPage,
		Duration:      activity.Duration,
		PageVisitedAt: timestamppb.New(activity.PageVisitedAt),
	}
}

func convertActivities(activities []db.Activity) []*pb.Activity {
	res := []*pb.Activity{}
	for i := 0; i < len(activities); i++ {
		res = append(res, &pb.Activity{
			Id:            activities[i].ID.String(),
			UserId:        activities[i].UserID.String(),
			ViewPage:      activities[i].ViewPage,
			Duration:      activities[i].Duration,
			PageVisitedAt: timestamppb.New(activities[i].PageVisitedAt),
		})
	}
	return res
}

func convertFavMovies(favMovies db.FavMovie) *pb.FavMovies {
	return &pb.FavMovies{
		UserId:   favMovies.UserID.String(),
		MovieIds: strings.Split(favMovies.Movies, " "),
	}
}

func convertMovie(movie mongodb.Movie) *pb.Movie {
	return &pb.Movie{
		Id:         movie.ID,
		Title:      movie.Title,
		Characters: movie.Characters,
		Genres:     movie.Genres,
		Rating:     movie.Rating,
	}
}

func convertMovies(movies []mongodb.Movie) []*pb.Movie {
	res := []*pb.Movie{}
	for i := 0; i < len(movies); i++ {
		res = append(res, &pb.Movie{
			Id:         movies[i].ID,
			Title:      movies[i].Title,
			Characters: movies[i].Characters,
			Genres:     movies[i].Genres,
			Rating:     movies[i].Rating,
		})
	}
	return res
}

func convertToken(accessToken string, expiresAt time.Time) *pb.Token {
	return &pb.Token{
		Token:     accessToken,
		ExpiresAt: timestamppb.New(expiresAt),
	}
}
