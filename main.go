package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/basic/docs"

	"ribeirosaimon/gobooplay/routers"
)

func main() {
	docs.SwaggerInfo.Title = "GoBoo Play"
	docs.SwaggerInfo.Description = "GoBoo Play challange with Golang"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r := gin.Default()

	routers.CreateConfigRouter(r)

	r.Run()
}
