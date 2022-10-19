package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type accountController struct {
	service accountService
}

func controller() accountController {
	return accountController{service: service()}
}

func (controller accountController) findAccount(context *gin.Context) {
	controller.service.saveAccount()
	context.JSON(http.StatusOK, gin.H{
		"message": "teste",
	})
}
