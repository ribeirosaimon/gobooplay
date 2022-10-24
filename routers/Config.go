package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"ribeirosaimon/gobooplay/api/account"
)

func CreateConfigRouter(e *gin.Engine) {
	e.GET("swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))

	version := e.Group("/v1")
	account.RouteAccount(version)
}
