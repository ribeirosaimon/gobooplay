package util

import (
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"ribeirosaimon/gobooplay/domain"
	"time"
)

func GetInitialProductId() primitive.ObjectID {
	hex, err := primitive.ObjectIDFromHex("000000000000000000000000")
	if err != nil {
		panic(err)
	}
	return hex
}

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
