package product

import (
	"github.com/gin-gonic/gin"
	"ribeirosaimon/gobooplay/middleware"
)

func ProductRoute(e *gin.RouterGroup) {
	group := e.Group("/product")
	group.Use(middleware.Authorization([]string{"ADMIN"})).GET("/teste", productController)
}
