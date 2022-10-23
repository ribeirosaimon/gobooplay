package account

import (
	"github.com/gin-gonic/gin"
)

func RouteAccount(e *gin.RouterGroup) {
	group := e.Group("/account")
	group.POST("/signup", controller().signUp)
	group.POST("/login", controller().login)

}
