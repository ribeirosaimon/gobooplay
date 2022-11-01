package routers

import (
	"github.com/gin-gonic/gin"
	"ribeirosaimon/gobooplay/api/account"
	"ribeirosaimon/gobooplay/api/product"
	"ribeirosaimon/gobooplay/api/shoppingCart"
)

func CreateConfigRouter(e *gin.Engine) {
	version := e.Group("api/v1")
	account.RouteAccount(version)
	product.RouteProduct(version)
	shoppingCart.RouteShoppingCart(version)
}
