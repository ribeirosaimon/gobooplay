package security

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/repository"
	"time"
)

var secretKey = ""

func init() {
	key := make([]byte, 64)
	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	secretKey = base64.StdEncoding.EncodeToString(key)
}

func CreateToken(account domain.Account) (string, error) {
	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Hour * 24).Unix()
	permission["userId"] = account.ID.Hex()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, permission).SignedString([]byte(secretKey))
}

func ValidationToken(token string) (domain.LoggedUser, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFunc()

	parseToken, err := jwt.Parse(token, verifyKey)
	if err != nil {
		return domain.LoggedUser{}, err
	}
	claims, ok := parseToken.Claims.(jwt.MapClaims)

	if ok && parseToken.Valid {
		userId := claims["userId"]
		rep := repository.NewAccountRepository()
		userDb, err := rep.FindById(ctx, fmt.Sprint(userId))
		if err != nil {
			return domain.LoggedUser{}, err
		}
		return domain.LoggedUser{
			Login:  userDb.Login,
			UserId: userDb.ID.Hex(),
			Role:   userDb.Role,
		}, nil
	}

	return domain.LoggedUser{}, errors.New("invalid Token")
}

func verifyKey(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("erro in token method")
	}
	return []byte(secretKey), nil
}
