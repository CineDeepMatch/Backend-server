package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	Querier
}
type MongoDBStore struct {
	*Queries
}

func NewStore(connMongoDB *mongo.Client, db string, collection string) Store {
	return &MongoDBStore{
		Queries: New(connMongoDB, db, collection),
	}
}
