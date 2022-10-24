package exceptions

import (
	"github.com/gin-gonic/gin"
	"time"
)

type HttpResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func ValidateException(c *gin.Context, msg string, status int) {
	c.AbortWithStatusJSON(status, HttpResponse{Message: msg, Timestamp: time.Now()})
}
