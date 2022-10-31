package product

import (
	"github.com/gin-gonic/gin"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/middleware"
)

func RouteProduct(e *gin.RouterGroup) {
	group := e.Group("/product")

	group.Use(middleware.Authorization([]string{domain.ADMIN, domain.GENERAL_ADMIN})).POST("/", ControllerProduct().SaveProduct)
	group.Use(middleware.Authorization([]string{domain.ADMIN, domain.GENERAL_ADMIN})).DELETE("/:productId", ControllerProduct().DeleteProduct)
	group.Use(middleware.Authorization([]string{domain.ADMIN, domain.GENERAL_ADMIN})).PUT("/:productId", ControllerProduct().UpdateProduct)
	group.Use(middleware.Authorization([]string{domain.ADMIN, domain.GENERAL_ADMIN})).GET("/available-subscribe", ControllerProduct().FindAvailableProduct)
}
