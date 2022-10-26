package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"ribeirosaimon/gobooplay/api/repository/mongoInterface"
	"ribeirosaimon/gobooplay/config"
	"ribeirosaimon/gobooplay/domain"
)

const collectionProduct = "product"

type Product struct {
	database *mongo.Collection
}

func NewProductRepository() mongoInterface.Product {
	return Product{database: config.DbConnection(collectionProduct)}
}

func (p Product) Save(ctx context.Context, product *domain.Product) (domain.Product, error) {
	panic("implement me")
}

func (p Product) FindById(ctx context.Context, s string) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p Product) DelebeById(ctx context.Context, s string) error {
	//TODO implement me
	panic("implement me")
}
