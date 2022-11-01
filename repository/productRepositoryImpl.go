package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"ribeirosaimon/gobooplay/config"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/repository/mongoInterface"
	"time"
)

const collectionProduct = "product"

type Product struct {
	database *mongo.Collection
}

func NewProductRepository() mongoInterface.Product {
	return Product{database: config.DbConnection(collectionProduct)}
}

func (p Product) Save(ctx context.Context, product *domain.Product) (domain.Product, error) {
	if product.ID != primitive.NilObjectID {
		update := bson.D{
			{"$set", bson.D{
				{"name", product.Name},
				{"price", product.Price},
				{"subscriptionTime", product.SubscriptionTime},
				{"updatedAt", time.Now()}}},
		}
		id, err := p.database.UpdateByID(ctx, product.ID, update)
		if err != nil || id.UpsertedCount != 1 {
			return domain.Product{}, err
		}
		byId, err := p.FindById(ctx, product.ID.String())
		if err != nil {
			return domain.Product{}, err
		}
		return byId, nil
	}
	one, err := p.database.InsertOne(ctx, product)
	if err != nil {
		return domain.Product{}, err
	}
	product.ID = one.InsertedID.(primitive.ObjectID)
	return *product, nil
}

func (p Product) FindById(ctx context.Context, id string) (domain.Product, error) {
	var product domain.Product
	filter := bson.D{
		{"_id", id},
	}
	err := p.database.FindOne(ctx, filter).Decode(&product)

	if err != nil {
		return product, err
	}
	return product, nil
}

func (p Product) DelebeById(ctx context.Context, id string) error {
	filter := bson.D{{"_id", id}}
	documents, err := p.database.CountDocuments(ctx, filter)
	if err != nil || documents != 1 {
		return err
	}
	p.database.FindOneAndDelete(ctx, filter)
	return nil
}
