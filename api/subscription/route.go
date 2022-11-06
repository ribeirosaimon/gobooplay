package subscription

import (
	"github.com/gin-gonic/gin"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/middleware"
)

func RouteSubscription(e *gin.RouterGroup) {
	group := e.Group("/subscription")
	group.Use(middleware.Authorization([]domain.Role{domain.USER})).GET("/", ControllerProduct().findMySubscription)
	group.Use(middleware.Authorization([]domain.Role{domain.USER})).POST("/pause", ControllerProduct().pauseSubscription)
	group.Use(middleware.Authorization([]domain.Role{domain.USER})).GET("/validate", ControllerProduct().validationSubscription)
}
