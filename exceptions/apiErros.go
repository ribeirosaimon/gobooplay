package exceptions

import (
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

type HttpResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func ValidateException(c *gin.Context, msg string, status int) {
	c.Error(errors.New(msg))
	c.JSON(status, HttpResponse{Message: msg, Timestamp: time.Now()})
}
