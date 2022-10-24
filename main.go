package main

import (
	"github.com/gin-gonic/gin"
	"ribeirosaimon/gobooplay/routers"
)

func main() {
	r := gin.Default()

	routers.CreateConfigRouter(r)

	r.Run(":8080")
}
