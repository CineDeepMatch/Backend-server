package mongodb

import (
	"context"
)

type Querier interface {
	CreateMovie(ctx context.Context, doc *Movie) (Movie, error)
	CreateManyMovies(ctx context.Context, doc *[]Movie) error
	GetMovieById(ctx context.Context, movieId string) (Movie, error)
}

var _ Querier = (*Queries)(nil)
