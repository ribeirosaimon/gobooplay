package order

import (
	"github.com/gin-gonic/gin"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/repository"
)

type orderService struct {
	productRepository      repository.MongoTemplateStruct[domain.Product]
	shoopingCartRepository repository.MongoTemplateStruct[domain.ShoppingCart]
}

func ServiceShoop() orderService {
	return orderService{
		productRepository:      repository.MongoTemplate[domain.Product](),
		shoopingCartRepository: repository.MongoTemplate[domain.ShoppingCart](),
	}
}

func (s shoopService) sendOrder(c *gin.Context, user domain.LoggedUser) {
	user.
}
