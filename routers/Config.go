package routers

import (
	"github.com/gin-gonic/gin"
	"ribeirosaimon/gobooplay/api/account"
)

func CreateConfigRouter(e *gin.Engine) {
	version := e.Group("/v1")
	account.RouteAccount(version)
}
