package mongodb

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImportCSV(t *testing.T) {
	f, err := os.Open("../data/movies.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	movies := ImportMovieCSV(data)
	movies_slice := movies[1:100]

	err = testStore.CreateManyMovies(context.Background(), &movies_slice)

	require.NoError(t, err)

}
