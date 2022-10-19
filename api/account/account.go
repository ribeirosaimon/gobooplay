package account

import (
	"github.com/gin-gonic/gin"
)

func RouteAccount(e *gin.Engine) {
	e.GET("/", controller().findAccount)
}
