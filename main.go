package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/repository"
	"ribeirosaimon/gobooplay/routers"
	"time"
)

func init() {
	mongoTemplate := repository.MongoTemplate[domain.Product]()
	ctx := context.Background()

	decimal128, err := primitive.ParseDecimal128("0")
	if err != nil {
		panic(err)
	}
	var product = domain.Product{
		Name:             "trial",
		Status:           domain.TRIAL,
		Price:            decimal128,
		SubscriptionTime: 1,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	filter := bson.D{
		{"status", domain.TRIAL},
	}
	exist, err := mongoTemplate.CountWithFilter(ctx, filter)
	if err != nil {
		panic(err)
	}
	if exist == 0 {
		mongoTemplate.Save(ctx, product)
	}
}

func main() {
	r := gin.Default()
	routers.CreateConfigRouter(r)
	r.Run(":8080")
}
