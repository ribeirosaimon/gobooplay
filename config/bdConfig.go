package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var clientInstance *mongo.Client
var mongoOnce sync.Once

const (
	//CONNECTIONSTRING = "mongodb://saimon:frajolinha202@localhost:27017"
	CONNECTIONSTRING = "mongodb://localhost:27017"
	DB               = "gobooplay"
)

func DbConnection(collection string) *mongo.Collection {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			panic(err)
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			panic(err)
		}
		clientInstance = client
	})
	return clientInstance.Database(DB).Collection(collection)
}
