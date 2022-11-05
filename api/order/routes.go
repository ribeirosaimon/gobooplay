package order

import (
	"github.com/gin-gonic/gin"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/middleware"
)

func RouteOrder(e *gin.RouterGroup) {
	group := e.Group("/order")
	group.Use(middleware.Authorization([]domain.Role{domain.USER})).POST("/send", ControllerShoop().SendOrder)
}
