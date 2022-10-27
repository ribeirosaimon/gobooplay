package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"ribeirosaimon/gobooplay/api/repository/mongoInterface"
	"ribeirosaimon/gobooplay/config"
)

type MongoTemplate struct {
	database *mongo.Collection
	myStruct interface{}
}

func (m MongoTemplate) findById(ctx context.Context, s string) (T, error) {
	//TODO implement me
	panic("implement me")
}

func NewMongoTemplateRepository[T any]() mongoInterface.GenericMongoTemplate[T] {
	return MongoTemplate{database: config.DbConnection(reflect.TypeOf(T).Name()), myStruct: T}
}
