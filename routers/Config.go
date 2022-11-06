package routers

import (
	"github.com/gin-gonic/gin"
	"ribeirosaimon/gobooplay/api/account"
	"ribeirosaimon/gobooplay/api/movies"
	"ribeirosaimon/gobooplay/api/order"
	"ribeirosaimon/gobooplay/api/product"
	"ribeirosaimon/gobooplay/api/shoppingCart"
	"ribeirosaimon/gobooplay/api/subscription"
	"ribeirosaimon/gobooplay/api/voucher"
)

func CreateConfigRouter(e *gin.Engine) {
	version := e.Group("api/v1")
	account.RouteAccount(version)
	product.RouteProduct(version)
	shoppingCart.RouteShoppingCart(version)
	order.RouteOrder(version)
	subscription.RouteSubscription(version)
	movies.RouteMovie(version)
	voucher.RouteVoucher(version)
}
