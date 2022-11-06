package movies

import (
	"github.com/gin-gonic/gin"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/middleware"
)

func RouteMovie(e *gin.RouterGroup) {
	group := e.Group("/movies")
	group.Use(middleware.Authorization([]domain.Role{domain.USER})).GET("/", ControllerMovie().GetMovie)
}
