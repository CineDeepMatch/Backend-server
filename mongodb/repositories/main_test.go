package mongodb

import (
	"context"
	"log"
	"os"
	"testing"

	util "github.com/CineDeepMatch/Backend-server/db/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
		return
	}

	connMongoDB, err := mongo.Connect(context.TODO(), options.Client().ApplyURI((config.MongoDBSource)))
	if err != nil {
		log.Fatal("cannot connect to mongodb:", err)
		return
	}

	defer connMongoDB.Disconnect(context.Background())

	testStore = NewStore(connMongoDB, config.MongoDB_DBTest, config.MongoDB_CollectionTest)

	os.Exit(m.Run())
}
