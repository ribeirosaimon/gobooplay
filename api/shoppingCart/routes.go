package shoppingCart

import (
	"github.com/gin-gonic/gin"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/middleware"
)

func RouteShoppingCart(e *gin.RouterGroup) {
	group := e.Group("/shopping-cart")
	group.Use(middleware.Authorization([]domain.Role{domain.USER})).GET("/", ControllerProduct().GetMyShoppingCart)
	group.Use(middleware.Authorization([]domain.Role{domain.USER})).POST("product/:productId", ControllerProduct().SaveProductInShoppingCart)
	group.Use(middleware.Authorization([]domain.Role{domain.USER})).POST("voucher/:voucherId", ControllerProduct().SaveVoucherInShoppingCart)
	group.Use(middleware.Authorization([]domain.Role{domain.USER})).POST("/clear", ControllerProduct().ClearShoppingCart)
}
