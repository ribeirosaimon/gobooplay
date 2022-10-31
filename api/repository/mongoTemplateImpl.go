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
	idReturned, err := m.ExistAndReturnId(ctx, id)
	if err != nil {
		return m.myInterface, errors.New("this document not exist")
	}

	update := bson.D{{"$set", filter}}

	updatedValue, err := m.database.UpdateOne(ctx, bson.D{{"_id", idReturned}}, update)

	if err != nil || updatedValue.MatchedCount == 0 {
		return m.myInterface, errors.New("something is wrong")
	}
	dbUpdated, err := m.FindById(ctx, idReturned.Hex())
	if err != nil {
		return m.myInterface, err
	}
	return dbUpdated, nil
}
func (m MongoTemplateStruct[T]) DeleteById(ctx context.Context, id string) error {
	idReturned, err := m.ExistAndReturnId(ctx, id)

	if err != nil {
		return errors.New("this document does not exists")
	}

	filter := bson.D{{"_id", idReturned}}

	one, err := m.database.DeleteOne(ctx, filter)
	if err != nil || one.DeletedCount != 1 {
		return err
	}
	return nil

}

func (m MongoTemplateStruct[T]) ExistAndReturnId(ctx context.Context, id string) (primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	filter := bson.D{{"_id", objectId}}
	documents, err := m.database.CountDocuments(ctx, filter)
	if err != nil || documents != 1 {
		return primitive.ObjectID{}, err
	}
	return objectId, nil
}
