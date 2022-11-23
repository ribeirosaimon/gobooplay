package shoppingCart

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/repository"
	"ribeirosaimon/gobooplay/util"
	"time"
)

type shoppingCartService struct {
	shoppingCartRepository repository.MongoTemplateStruct[domain.ShoppingCart]
	productRepository      repository.MongoTemplateStruct[domain.Product]
	userRepository         repository.MongoTemplateStruct[domain.Account]
	voucherRepository      repository.MongoTemplateStruct[domain.Voucher]
}

func ServiceShoppingCart() shoppingCartService {
	return shoppingCartService{
		shoppingCartRepository: repository.MongoTemplate[domain.ShoppingCart](),
		productRepository:      repository.MongoTemplate[domain.Product](),
		userRepository:         repository.MongoTemplate[domain.Account](),
		voucherRepository:      repository.MongoTemplate[domain.Voucher](),
	}
}

func (s shoppingCartService) saveProductShoppingCart(c context.Context, id string, loggedUser domain.LoggedUser) (domain.ShoppingCart, error) {
	//if !util.ContainsRole[domain.Role](loggedUser.Role, domain.USER) {
	//	return domain.ShoppingCart{}, errors.New("you not have permission")
	//}

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
	newShoppingCart.FinalPrice = productDb.Price

	savedShoppingCart, err := s.shoppingCartRepository.Save(c, newShoppingCart)
	if err != nil {
		return domain.ShoppingCart{}, err
	}

	return savedShoppingCart, nil
}

func (s shoppingCartService) saveVoucherShoppingCart(
	c *gin.Context, voucheId string, loggedUser domain.LoggedUser) (domain.ShoppingCart, error) {

	user, err := s.userRepository.FindById(c, loggedUser.UserId)
	if err != nil {
		return domain.ShoppingCart{}, err
	}
	shoppingCartFilter := bson.D{
		{"owner.userId", user.ID.Hex()},
	}

	shoppingCart, err := s.shoppingCartRepository.FindOneByFilter(c, shoppingCartFilter)
	if err != nil || &shoppingCart.Product == nil {
		return domain.ShoppingCart{}, errors.New("you not have a Shopping Cart")
	}

	voucherDb, err := s.voucherRepository.FindById(c, voucheId)
	if err != nil {
		return domain.ShoppingCart{}, errors.New("product not found")
	}

	if voucherDb.Quantity == 0 {
		return domain.ShoppingCart{}, errors.New("voucher is finish")
	}

	shoppingCart.Voucher = voucherDb

	promotionValue, err := decimal.NewFromString(shoppingCart.Price.String())
	voucherValue, err := decimal.NewFromString(voucherDb.Price.String())
	if err != nil {
		return domain.ShoppingCart{}, err
	}

	promotion := promotionValue.Sub(voucherValue)
	if promotion.IsNegative() {
		decimal128, err := primitive.ParseDecimal128("0")
		if err != nil {
			return domain.ShoppingCart{}, err
		}

		shoppingCart.FinalPrice = decimal128
	} else {
		decimal128, err := primitive.ParseDecimal128(promotion.String())
		if err != nil {
			return domain.ShoppingCart{}, err
		}

		shoppingCart.FinalPrice = decimal128
	}
	return shoppingCart, nil
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
