package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ribeirosaimon/gobooplay/api/product"
	"ribeirosaimon/gobooplay/api/subscription"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/repository"
	"ribeirosaimon/gobooplay/routers"
	"ribeirosaimon/gobooplay/security"
	"time"
)

func init() {
	mongoTemplate := repository.MongoTemplate[domain.Product]()
	ctx := context.Background()

	decimal128, err := primitive.ParseDecimal128("0")
	if err != nil {
		panic(err)
	}
	var pdt = domain.Product{
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
		mongoTemplate.Save(ctx, pdt)
	}

	mongoTemplateAccount := repository.MongoTemplate[domain.Account]()
	accountFilter := bson.D{
		{"role", domain.ADMIN},
	}
	countAdmin, err := mongoTemplateAccount.CountWithFilter(ctx, accountFilter)
	if err != nil {
		panic(err)
	}
	if countAdmin == 0 {
		password, _ := security.EncriptyPassword("admin")
		var acc = domain.Account{
			Name:     "admin",
			Password: string(password),
			Login:    "admin",
			Status:   domain.ACTIVE,
		}
		mongoTemplateAccount.Save(ctx, acc)
		product.ServiceProduct().GetTrialProduct(ctx)
		firstProduct, err := product.ServiceProduct().GetTrialProduct(ctx)
		if err != nil {
			panic(err)
		}
		subscription.ServiceSubscription().CreateSubscription(ctx, acc, firstProduct)
	}
}

func main() {
	r := gin.Default()
	routers.CreateConfigRouter(r)
	r.Run(":8080")
}
