package mongodb

import (
	"context"
	"testing"

	util "github.com/CineDeepMatch/Backend-server/db/utils"
	"github.com/stretchr/testify/require"
)

func createRandomGenres(t *testing.T, n int) []string {
	var genres []string
	for i := 0; i < n; i++ {
		var genre = util.RandomString(6)
		genres = append(genres, genre)
	}
	return genres
}

func createRandomCharacters(t *testing.T, n int) []string {
	var characters []string
	for i := 0; i < n; i++ {
		var character = util.RandomString(6)
		characters = append(characters, character)
	}
	return characters
}

func createRandomMovie(t *testing.T) *Movie {
	return &Movie{
		ID:         util.RandomString(15),
		Title:      util.RandomString(15),
		Genres:     createRandomGenres(t, 5),
		Characters: createRandomCharacters(t, 5),
	}
}

func TestCreateMovie(t *testing.T) {
	movie := createRandomMovie(t)

	movie_create, err := testStore.CreateMovie(context.Background(), movie)

	require.NoError(t, err)
	require.Equal(t, movie.ID, movie_create.ID)
	require.Equal(t, movie.Title, movie_create.Title)
	require.Equal(t, movie.Genres, movie_create.Genres)
	require.Equal(t, movie.Characters, movie_create.Characters)
}

func TestGetMovie(t *testing.T) {
	movie := createRandomMovie(t)

	testStore.CreateMovie(context.Background(), movie)

	movie_get, err := testStore.GetMovieById(context.Background(), movie.ID)

	require.NoError(t, err)
	require.Equal(t, movie.ID, movie_get.ID)
	require.Equal(t, movie.Title, movie_get.Title)
	require.Equal(t, movie.Genres, movie_get.Genres)
	require.Equal(t, movie.Characters, movie_get.Characters)

}
