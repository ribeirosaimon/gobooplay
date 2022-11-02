package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/repository"
	"ribeirosaimon/gobooplay/routers"
	"ribeirosaimon/gobooplay/util"
	"time"
)

func createInitialProduct() {

	decimal128, err := primitive.ParseDecimal128("0")
	if err != nil {
		panic(err)
	}
	var product = domain.Product{
		ID:               util.GetInitialProductId(),
		Name:             "trial",
		Price:            decimal128,
		SubscriptionTime: 1,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	repository.MongoTemplate[domain.Product]().Save(context.Background(), product)
}

func main() {
	createInitialProduct()

	r := gin.Default()
	routers.CreateConfigRouter(r)
	r.Run(":8080")
}
