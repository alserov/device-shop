package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var (
	key = os.Getenv("SECRET_KEY")
)

func ValidateToken(t string) error {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return err
	}

	if int64(token.Claims.(jwt.MapClaims)["exp"].(float64)) < time.Now().Local().Unix() {
		return errors.New("token expired")
	}

	return nil
}

func CheckIfAdmin(t string) error {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return err
	}

	exp := token.Claims.(jwt.MapClaims)["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		return errors.New("token expired")
	}

	if token.Claims.(jwt.MapClaims)["role"] != "admin" {
		return errors.New("not allowed")
	}

	return nil
}
