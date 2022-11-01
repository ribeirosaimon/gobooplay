package util

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"ribeirosaimon/gobooplay/domain"
	"strings"
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
	hmd5 := md5.Sum([]byte(time.Now().String()))
	stringValue := fmt.Sprintf("%x", hmd5)[:5]
	var stringResponse = ""
	for x := range stringValue {
		var y string
		s := string(rune(x))
		int1 := rand.Intn(1000)
		int2 := rand.Intn(1000)
		if int1 >= int2 {
			y = strings.ToUpper(s)
		}
		fmt.Println(s)
		y = s
		stringResponse += y
	}
	return stringResponse
}
