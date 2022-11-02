package order

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/repository"
)

type orderService struct {
	productRepository      repository.MongoTemplateStruct[domain.Product]
	shoopingCartRepository repository.MongoTemplateStruct[domain.ShoppingCart]
	userRepository         repository.MongoTemplateStruct[domain.Account]
}

func ServiceShoop() orderService {
	return orderService{
		productRepository:      repository.MongoTemplate[domain.Product](),
		shoopingCartRepository: repository.MongoTemplate[domain.ShoppingCart](),
		userRepository:         repository.MongoTemplate[domain.Account](),
	}
}

func (s orderService) sendOrder(c *gin.Context, loggedUser domain.LoggedUser) (domain.Account, error) {
	user, err := s.userRepository.FindById(c, loggedUser.UserId)
	if err != nil {
		return domain.Account{}, err
	}

	userFilter := bson.D{
		{"owner.userId", user.ID},
	}
	shoppingCart, err := s.shoopingCartRepository.FindOneByFilter(c, userFilter)
	if err != nil {
		return domain.Account{}, err
	}

	user
}
