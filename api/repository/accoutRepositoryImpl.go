package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"ribeirosaimon/gobooplay/api/repository/mongoInterface"
	"ribeirosaimon/gobooplay/config"
	"ribeirosaimon/gobooplay/domain"
	"time"
)

const collectionAccount = "Account"

type Account struct {
	database *mongo.Collection
}

func NewAccountRepository() mongoInterface.Account {
	return Account{database: config.DbConnection(collectionAccount)}
}

func (conn Account) Save(context context.Context, account *domain.Account) (domain.Account, error) {
	existUser := conn.ExistUserWithLogin(context, account.Login)
	if existUser {
		update := bson.D{
			{"$set", bson.D{
				{"name", account.Name},
				{"password", account.Password},
				{"familyName", account.FamilyName},
				{"login", account.Login},
				{"gender", account.Gender},
				{"role", account.Role},
				{"lastLogin", account.LastLogin},
				{"lastLoginAttemp", account.LastLoginAttemp},
				{"passwordErrorCount", account.PasswordErrorCount},
				{"loginCount", account.LoginCount},
				{"updatedAt", time.Now()}}},
		}

		id, err := conn.database.UpdateByID(context, account.ID, update)
		if err != nil {

			return domain.Account{}, err
		}
		if id.ModifiedCount != 0 {
			return *account, nil
		}
	} else {
		one, err := conn.database.InsertOne(context, account)
		if err != nil {
			return domain.Account{}, err
		}

		account.ID = one.InsertedID.(primitive.ObjectID)
		return *account, nil
	}
	return domain.Account{}, nil
}

func (conn Account) FindById(context context.Context, id string) (domain.Account, error) {
	var account domain.Account

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return account, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}

	if err := conn.database.FindOne(context, filter).Decode(&account); err != nil {
		return account, err
	}
	return account, nil
}

func (conn Account) FindAccountByLogin(context context.Context, login string) (domain.Account, error) {
	var account domain.Account

	filter := bson.D{primitive.E{Key: "login", Value: login}}
	err := conn.database.FindOne(context, filter).Decode(&account)

	if err != nil {
		return domain.Account{}, err
	}

	return account, nil
}

func (conn Account) DelebeById(ctx context.Context, s string) error {
	//TODO implement me
	panic("implement me")
}

func (conn Account) ExistUserWithLogin(ctx context.Context, login string) bool {
	filter := bson.D{primitive.E{Key: "login", Value: login}}
	documents, err := conn.database.CountDocuments(ctx, filter)
	if err != nil || documents == 0 {
		return false
	}
	return true
}
