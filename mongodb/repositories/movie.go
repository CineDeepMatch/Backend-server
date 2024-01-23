package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
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
