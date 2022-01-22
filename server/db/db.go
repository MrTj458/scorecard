package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx = context.Background()

// Connect connects to the MongoDB database with the uri and name provided,
// panics if the connection can't be made
func Connect(uri, dbName string) *mongo.Database {
	client, err := mongo.Connect(Ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(dbName)
}
