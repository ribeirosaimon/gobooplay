package shoppingCart

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/repository"
	"ribeirosaimon/gobooplay/util"
	"time"
)

type shoppingCartService struct {
	shoppingCartRepository repository.MongoTemplateStruct[domain.ShoppingCart]
	productRepository      repository.MongoTemplateStruct[domain.Product]
	userRepository         repository.MongoTemplateStruct[domain.Account]
}

func ServiceShoppingCart() shoppingCartService {
	return shoppingCartService{
		shoppingCartRepository: repository.MongoTemplate[domain.ShoppingCart](),
		productRepository:      repository.MongoTemplate[domain.Product](),
		userRepository:         repository.MongoTemplate[domain.Account](),
	}
}

func (s shoppingCartService) saveShoppingCart(c context.Context, id string, loggedUser domain.LoggedUser) (domain.ShoppingCart, error) {
	if !util.ContainsRole[string](loggedUser.Role, domain.USER) {
		return domain.ShoppingCart{}, errors.New("you not have permission")
	}

	user, err := s.userRepository.FindById(c, loggedUser.UserId)
	if err != nil {
		return domain.ShoppingCart{}, err
	}

	shoppingCartFilter := bson.D{
		{"owner.userId", id},
	}

	shoppingCart, err := s.shoppingCartRepository.Find(c, shoppingCartFilter)
	if len(shoppingCart) != 0 {
		return domain.ShoppingCart{}, errors.New("you have a Shopping Cart")
	}
	var newShoppingCart domain.ShoppingCart

	newShoppingCart.Owner = user.MyRef()
	newShoppingCart.Hash = util.CreateHash()
	fmt.Println(util.CreateHash())
	newShoppingCart.CreatedAt = time.Now()
	newShoppingCart.UpdateAt = time.Now()

	find, err := s.productRepository.FindById(c, id)
	if err != nil {
		return domain.ShoppingCart{}, err
	}
	fmt.Println(find)
	return domain.ShoppingCart{}, nil
}
