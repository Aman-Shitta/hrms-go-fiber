package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var MGinstance *MongoInstance

const dbName = "test" // "hrms-go-fiber"
const mongoURI = "mongodb://localhost:27017/" + dbName

func NewMongoInstance() error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}

	MGinstance = &MongoInstance{
		Client: client,
		Db:     client.Database(dbName),
	}
	return nil

}
