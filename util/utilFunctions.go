package util

import (
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"ribeirosaimon/gobooplay/domain"
	"time"
)

func GetUser(c *gin.Context) domain.LoggedUser {
	loggedUser, _ := c.Get("loggedUser")
	return loggedUser.(domain.LoggedUser)
}

func ContainsRole[T comparable](sliceRoles []T, roleComparable T) bool {
	for _, v := range sliceRoles {
		if v == roleComparable {
			return true
		}
	}
	return false
}

func CreateHash() string {
	rand.Seed(time.Now().UnixNano())
	stringValue := sha256.Sum256([]byte(string(rand.Intn(1000))))
	return fmt.Sprintf("%x", stringValue)[:5]
}
