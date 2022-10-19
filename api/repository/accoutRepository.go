package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"ribeirosaimon/gobooplay/config"
	"ribeirosaimon/gobooplay/domain"
)

const collectionAccount = "Account"

type Account struct {
	mongo *mongo.Collection
}

func NewAccountRepository() Account {
	return Account{mongo: config.DbConnection(collectionAccount)}
}

func (conn Account) SaveAccount(context context.Context, account domain.Account) {

	conn.mongo.InsertOne(context)
}

func (conn Account) FindAccountById(context context.Context, id string) (domain.Account, error) {
	var account domain.Account

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return account, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}

	if err := conn.mongo.FindOne(context, filter).Decode(&account); err != nil {
		return account, err
	}
	return account, nil
}
