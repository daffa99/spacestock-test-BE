package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect func to connect into mongo database (Mongo Atlas)
func Connect() (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://daffa99:rahasia@cluster0-shard-00-00-hk5xc.mongodb.net:27017,cluster0-shard-00-01-hk5xc.mongodb.net:27017,cluster0-shard-00-02-hk5xc.mongodb.net:27017/test?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true&w=majority")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// Check testing env
	testEnv := os.Getenv("IS_TESTING")
	if testEnv == "true" {
		return client.Database("spacestock_be_test"), nil
	}
	return client.Database("spacestock_be"), nil
}
