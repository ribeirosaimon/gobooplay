package security

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func EncriptyPassword(pass string) ([]byte, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return []byte{}, errors.New("error encripty password")
	}
	return password, nil
}

func VerifyPassword(hashPass, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))
}
