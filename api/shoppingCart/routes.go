package shoppingCart

import (
	"github.com/gin-gonic/gin"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/middleware"
)

func RouteShoppingCart(e *gin.RouterGroup) {
	group := e.Group("/shopping-cart")
	group.Use(middleware.Authorization([]string{domain.USER})).GET("/", ControllerProduct().GetMyShoppingCart)
	group.Use(middleware.Authorization([]string{domain.USER})).POST("/:productId", ControllerProduct().SaveShoppingCart)
	group.Use(middleware.Authorization([]string{domain.USER})).POST("/clear", ControllerProduct().ClearShoppingCart)
}
