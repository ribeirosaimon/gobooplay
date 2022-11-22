package shoppingCart

import (
	"context"
	"errors"
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
	if !util.ContainsRole[domain.Role](loggedUser.Role, domain.USER) {
		return domain.ShoppingCart{}, errors.New("you not have permission")
	}

	user, err := s.userRepository.FindById(c, loggedUser.UserId)
	if err != nil {
		return domain.ShoppingCart{}, err
	}

	shoppingCartFilter := bson.D{
		{"owner.userId", loggedUser.UserId},
	}

	shoppingCart, err := s.shoppingCartRepository.FindAllByFilter(c, shoppingCartFilter)
	if len(shoppingCart) == 1 {
		return domain.ShoppingCart{}, errors.New("you have a Shopping Cart")
	}

	productDb, err := s.productRepository.FindById(c, id)
	if err != nil {
		return domain.ShoppingCart{}, errors.New("product not found")
	}

	var newShoppingCart domain.ShoppingCart
	newShoppingCart.Owner = user.MyRef()
	newShoppingCart.Hash = util.CreateHash()
	newShoppingCart.CreatedAt = time.Now()
	newShoppingCart.UpdateAt = time.Now()
	newShoppingCart.Product = productDb
	newShoppingCart.Price = productDb.Price

	savedShoppingCart, err := s.shoppingCartRepository.Save(c, newShoppingCart)
	if err != nil {
		return domain.ShoppingCart{}, err
	}

	return savedShoppingCart, nil
}

func (s shoppingCartService) findShoppingCart(c context.Context, user domain.LoggedUser) (domain.ShoppingCart, error) {
	filter := bson.D{
		{"owner.userId", user.UserId},
	}
	cart, err := s.shoppingCartRepository.FindOneByFilter(c, filter)
	if err != nil {
		return domain.ShoppingCart{}, errors.New("you not have a shopping cart")
	}
	return cart, nil
}

func (s shoppingCartService) clearShoppingCart(c context.Context, user domain.LoggedUser) error {
	cart, err := s.findShoppingCart(c, user)
	if err != nil {
		return err
	}
	if err := s.shoppingCartRepository.DeleteById(c, cart.ID.Hex()); err != nil {
		return err
	}
	return nil
}
