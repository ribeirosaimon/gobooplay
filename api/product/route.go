package product

import (
	"github.com/gin-gonic/gin"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/middleware"
)

func RouteProduct(e *gin.RouterGroup) {
	group := e.Group("/product")

	group.Use(middleware.Authorization([]domain.Role{domain.ADMIN, domain.GENERAL_ADMIN, domain.USER})).
		GET("/available-subscribe", ControllerProduct().FindAvailableProduct)
	group.Use(middleware.Authorization([]domain.Role{domain.ADMIN, domain.GENERAL_ADMIN})).POST("/", ControllerProduct().saveProduct)
	group.Use(middleware.Authorization([]domain.Role{domain.ADMIN, domain.GENERAL_ADMIN})).DELETE("/:productId", ControllerProduct().deleteProduct)
	group.Use(middleware.Authorization([]domain.Role{domain.ADMIN, domain.GENERAL_ADMIN})).PUT("/:productId", ControllerProduct().updateProduct)
}
