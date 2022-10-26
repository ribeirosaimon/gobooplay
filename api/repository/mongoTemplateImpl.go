package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"ribeirosaimon/gobooplay/api/repository/mongoInterface"
	"ribeirosaimon/gobooplay/config"
)

type MongoTemplate struct {
	database *mongo.Collection
}

func NewMongoRepository() mongoInterface.MongoTemplate {
	return MongoTemplate{database: config.DbConnection(collectionProduct)}
}

func (m MongoTemplate) findById(ctx context.Context, s string, i interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}
