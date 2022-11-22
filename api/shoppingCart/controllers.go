package shoppingCart

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ribeirosaimon/gobooplay/exceptions"
	"ribeirosaimon/gobooplay/util"
)

type controllerShoppingCart struct {
	service shoppingCartService
}

func ControllerProduct() controllerShoppingCart {
	return controllerShoppingCart{
		service: ServiceShoppingCart(),
	}
}

func (s controllerShoppingCart) GetMyShoppingCart(c *gin.Context) {
	user := util.GetUser(c)
	cart, err := s.service.findShoppingCart(c, user)
	if err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}
	c.JSON(http.StatusCreated, cart)
	return
}

func (s controllerShoppingCart) SaveProductInShoppingCart(c *gin.Context) {
	productId := c.Param("productId")
	user := util.GetUser(c)
	cart, err := s.service.saveProductShoppingCart(c, productId, user)
	if err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}
	c.JSON(http.StatusCreated, cart)
	return
}

func (s controllerShoppingCart) SaveVoucherInShoppingCart(c *gin.Context) {
	voucherId := c.Param("voucherId")
	user := util.GetUser(c)
	cart, err := s.service.saveVoucherShoppingCart(c, voucherId, user)
	if err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}
	c.JSON(http.StatusCreated, cart)
	return
}

func (s controllerShoppingCart) ClearShoppingCart(c *gin.Context) {
	user := util.GetUser(c)

	if err := s.service.clearShoppingCart(c, user); err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}
	return
}
