package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (q *Queries) CreateMovie(ctx context.Context, doc *Movie) (Movie, error) {

	_, err := q.coll.InsertOne(context.TODO(), doc)

	return *doc, err
}

func (q *Queries) CreateManyMovies(ctx context.Context, doc *[]Movie) error {
	var newDoc []interface{}
	for _, movie := range *doc {
		newDoc = append(newDoc, movie)
	}
	_, err := q.coll.InsertMany(context.TODO(), newDoc)

	return err
}

func (q *Queries) GetMovieById(ctx context.Context, id string) (Movie, error) {
	var movie Movie
	filter := bson.M{"_id": id}
	err := q.coll.FindOne(context.TODO(), filter).Decode(&movie)
	return movie, err
}

func (q *Queries) GetManyMovies(ctx context.Context, pageNumber int64, pageSize int64) ([]Movie, error) {
	var movies []Movie

	skip := (pageNumber - 1) * pageSize

	findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(pageSize))

	cursor, err := q.coll.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		return movies, err
	}

	err = cursor.All(context.TODO(), &movies)

	return movies, err
}
