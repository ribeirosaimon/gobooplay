package order

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/repository"
	"time"
)

type OrderService struct {
	productRepository      repository.MongoTemplateStruct[domain.Product]
	shoopingCartRepository repository.MongoTemplateStruct[domain.ShoppingCart]
	userRepository         repository.MongoTemplateStruct[domain.Account]
	subscriptionRepository repository.MongoTemplateStruct[domain.Subscription]
}

func ServiceOrder() OrderService {
	return OrderService{
		productRepository:      repository.MongoTemplate[domain.Product](),
		shoopingCartRepository: repository.MongoTemplate[domain.ShoppingCart](),
		userRepository:         repository.MongoTemplate[domain.Account](),
		subscriptionRepository: repository.MongoTemplate[domain.Subscription](),
	}
}

func (s OrderService) sendOrder(c *gin.Context, loggedUser domain.LoggedUser) (domain.Subscription, error) {
	user, err := s.userRepository.FindById(c, loggedUser.UserId)
	if err != nil {
		return domain.Subscription{}, err
	}

	userFilter := bson.D{
		{"owner.userId", user.ID.Hex()},
	}
	shoppingCart, err := s.shoopingCartRepository.FindOneByFilter(c, userFilter)
	if err != nil {
		return domain.Subscription{}, errors.New("you not have shopping cart")
	}
	product, err := s.productRepository.FindById(c, shoppingCart.Product.ID.Hex())
	subsFilter := bson.D{
		{"owner.userId", user.ID.Hex()},
	}

	now := time.Now()

	mySubs, err := s.subscriptionRepository.FindOneByFilter(c, subsFilter)
	if err != nil {
		return domain.Subscription{}, err
	}

	filterSaved := bson.D{
		{"updatedAt", now},
		{"product", product},
		{"endAt", mySubs.EndAt.AddDate(0, int(product.SubscriptionTime), 0)},
	}
	mySubs, err = s.subscriptionRepository.UpdateById(c, mySubs.ID.Hex(), filterSaved)
	if err != nil {
		return domain.Subscription{}, err
	}

	if err := s.shoopingCartRepository.DeleteById(c, shoppingCart.ID.Hex()); err != nil {
		return domain.Subscription{}, err
	}
	return mySubs, nil
}
