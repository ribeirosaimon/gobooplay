package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"ribeirosaimon/gobooplay/config"
	"strings"
)

type MongoTemplateStruct[T any] struct {
	database        *mongo.Collection
	myInterface     T
	listOfInterface []T
}

func MongoTemplate[T interface{}]() MongoTemplateStruct[T] {
	var myInterface T
	return MongoTemplateStruct[T]{
		database:    config.DbConnection(reflect.TypeOf(myInterface).Name()),
		myInterface: myInterface,
	}
}

func (m MongoTemplateStruct[T]) FindById(ctx context.Context, s string) (T, error) {
	objectId, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		return m.myInterface, err
	}
	filter := bson.D{{"_id", objectId}}
	if err := m.database.FindOne(ctx, filter).Decode(&m.myInterface); err != nil {
		return m.myInterface, err
	}

	return m.myInterface, nil
}

func (m MongoTemplateStruct[T]) Find(ctx context.Context, filter bson.D) ([]T, error) {
	cur, err := m.database.Find(ctx, filter)
	defer cur.Close(ctx)
	if err != nil {
		return m.listOfInterface, err
	}

	for cur.Next(ctx) {

		if err := cur.Decode(&m.myInterface); err != nil {
			return m.listOfInterface, err
		}

		m.listOfInterface = append(m.listOfInterface, m.myInterface)
	}
	return m.listOfInterface, nil
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

func (m MongoTemplateStruct[T]) UpdateById(ctx context.Context, id string, filter bson.D) (T, error) {
	mapElement := filter.Map()
	for key, _ := range mapElement {
		if strings.Contains(key, "id") {
			return m.myInterface, errors.New("id canot be changed")
		}
	}
	haveValue, err := m.CountDocument(ctx, filter)
	if err != nil {
		return m.myInterface, err
	}
	if !haveValue {
		return m.myInterface, errors.New("this document not exist")
	}
	update := bson.D{{"$set", filter}}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return m.myInterface, err
	}

	updatedValue, err := m.database.UpdateByID(ctx, objectId, update)

	if updatedValue.ModifiedCount != 1 {
		return m.myInterface, errors.New("something is wrong")
	}
	dbUpdated, err := m.FindById(ctx, objectId.String())
	if err != nil {
		return m.myInterface, err
	}
	return dbUpdated, nil
}
func (m MongoTemplateStruct[T]) DeleteById(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", objectId}}

	document, err := m.CountDocument(ctx, filter)
	if err != nil {
		return err
	}
	one, err := m.database.DeleteOne(ctx, document)
	if err != nil || one.DeletedCount != 1 {
		return err
	}
	return nil

}

func (m MongoTemplateStruct[T]) CountDocument(ctx context.Context, filter bson.D) (bool, error) {
	documents, err := m.database.CountDocuments(ctx, filter)
	if err != nil || documents != 1 {
		return false, err
	}
	return true, nil
}
