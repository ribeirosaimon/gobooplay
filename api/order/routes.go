package order

import (
	"github.com/gin-gonic/gin"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/middleware"
)

func RouteShoop(e *gin.RouterGroup) {
	group := e.Group("/order")
	group.Use(middleware.Authorization([]string{domain.USER})).PUT("/send", ControllerShoop().SendOrder)
}
