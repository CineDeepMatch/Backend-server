package mongodb

import "go.mongodb.org/mongo-driver/mongo"

func New(connMongoDB *mongo.Client, db string, collection string) *Queries {
	return &Queries{coll: connMongoDB.Database(db).Collection(collection)}
}

type Queries struct {
	coll *mongo.Collection
}
