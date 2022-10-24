package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"ribeirosaimon/gobooplay/config"
)

const collectionProduct = "product"

type Product struct {
	database *mongo.Collection
}

func NewProductRepository() Product {
	return Product{database: config.DbConnection(_interface.collectionAccount)}
}
