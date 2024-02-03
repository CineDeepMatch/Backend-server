package mongodb

import (
	"context"
)

type Querier interface {
	CreateMovie(ctx context.Context, doc *Movie) (Movie, error)
	CreateManyMovies(ctx context.Context, doc *[]Movie) error
	GetMovieById(ctx context.Context, movieId string) (Movie, error)
	GetMovieByIds(ctx context.Context, ids []string) ([]Movie, error)
	GetManyMovies(ctx context.Context, pageNumber int64, pageSize int64) ([]Movie, error)
}

var _ Querier = (*Queries)(nil)
