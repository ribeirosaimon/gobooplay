package shoppingCart

import (
	"github.com/gin-gonic/gin"
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

}

func (s controllerShoppingCart) SaveShoppingCart(c *gin.Context) {
	productId := c.Param("productId")
	user := util.GetUser(c)
	s.service.saveShoppingCart(c, productId, user)
}
