package account

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/exceptions"
	"time"
)

type accountController struct {
	service accountService
}

func controller() accountController {
	return accountController{service: service()}
}

func (s accountController) signUp(c *gin.Context) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*80)
	defer cancelFunc()

	var payload domain.AccountDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&payload); err != nil {
		exceptions.ValidateException(c, "incorrect body", http.StatusConflict)
		return
	}

	account, err := s.service.saveAccount(ctx, payload)
	if err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}
	c.JSON(http.StatusCreated, account)
	return
}

func (s accountController) login(c *gin.Context) {

	var login domain.LoginDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&login); err != nil {
		exceptions.ValidateException(c, "incorrect body", http.StatusConflict)
		return
	}
	token, err := s.service.login(c, login)
	if err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}

	c.JSON(http.StatusOK, token)
	return
}
