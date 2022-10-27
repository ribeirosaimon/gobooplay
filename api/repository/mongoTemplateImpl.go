package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"ribeirosaimon/gobooplay/config"
)

type MongoTemplateStruct[T any] struct {
	database    *mongo.Collection
	myInterface T
}

func MongoTemplate[T interface{}]() MongoTemplateStruct[T] {
	var myInterface T
	return MongoTemplateStruct[T]{
		database:    config.DbConnection(reflect.TypeOf(myInterface).Name()),
		myInterface: myInterface,
	}
}

func (m MongoTemplateStruct[T]) FindById(ctx context.Context, s string) (T, error) {
	filter := bson.D{{"_id", s}}
	if err := m.database.FindOne(ctx, filter).Decode(&m.myInterface); err != nil {
		return m.myInterface, err
	}

	return m.myInterface, nil
}

func (m MongoTemplateStruct[T]) Save(ctx context.Context, myStruct T) (T, error) {
	value, err := m.database.InsertOne(ctx, myStruct)

	if err != nil {
		return m.myInterface, err
	}

	filter := bson.D{{"_id", value.InsertedID}}
	if err := m.database.FindOne(ctx, filter).Decode(&m.myInterface); err != nil {
		return m.myInterface, err
	}
	return m.myInterface, nil
}

func (m MongoTemplateStruct[T]) UpdateById(ctx context.Context, id string, filter bson.D) {
	if m.CountDocument(ctx, filter) {

	}

}

func (m MongoTemplateStruct[T]) CountDocument(ctx context.Context, filter bson.D) (bool, error) {
	documents, err := m.database.CountDocuments(ctx, filter)
	if err != nil || documents != 1 {
		return false, err
	}
	return true, nil
}
