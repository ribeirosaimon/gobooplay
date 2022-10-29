package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/exceptions"
	"ribeirosaimon/gobooplay/security"
	"strings"
	"time"
)

func Authorization(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		loggedUser, err := getToken(c)
		if err != nil {
			exceptions.ValidateException(c, "no have token", http.StatusConflict)
			return
		}
		authorization := contains(roles, loggedUser.Role)
		if !authorization {
			exceptions.ValidateException(c, "you not authorizated", http.StatusConflict)
			return
		}
		c.Set("loggedUser", loggedUser)
		latency := time.Since(t)
		log.Println(latency)
	}
}

func getToken(c *gin.Context) (domain.LoggedUser, error) {
	var token string
	headerToken := c.GetHeader("Authorization")

	if len(strings.Split(headerToken, " ")) == 2 {
		token = strings.Split(headerToken, " ")[1]
	} else {
		return domain.LoggedUser{}, errors.New("you need access token")
	}

	return security.ValidationToken(token)
}

func contains(loggedUserRole, routeRole []string) bool {
	for _, userRole := range loggedUserRole {

		for _, role := range routeRole {
			if userRole == role {
				return true
			}
		}
	}
	return false
}
