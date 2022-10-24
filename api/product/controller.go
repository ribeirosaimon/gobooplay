package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func productController(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
