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

// signUp             godoc
// @Summary      sign Up accounts
// @Description  Sign Up account.
// @Tags         signUp
// @Produce      json
// @Success      200  {object}  domain.AccountDTO
// @Router       /books [get]
func (s accountController) signUp(c *gin.Context) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*80)
	defer cancelFunc()

	var payload domain.AccountDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&payload); err != nil {
		exceptions.ValidateException(c, "incorrect body", http.StatusConflict)
		return
	}

	account, err := s.service.saveAccountService(ctx, payload)
	if err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}
	c.JSON(http.StatusCreated, account)
	return
}

// login godoc
// @Summary Login
// @ID      login
// @Produce json
// @Success 200 {object} domain.UserAccessToken
// @Failure 409 {object} exceptions.HttpResponse
// @Router  /account/login [POST]
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
