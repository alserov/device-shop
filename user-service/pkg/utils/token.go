package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var (
	key = os.Getenv("SECRET_KEY")
)

func GenerateTokens(role string) (string, string, error) {
	cl := jwt.MapClaims{}
	cl["role"] = role
	cl["exp"] = time.Now().Add(time.Hour * 164).Unix()

	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, cl).SignedString([]byte(key))
	if err != nil {
		return "", "", err
	}

	rClaims := jwt.MapClaims{}
	rClaims["exp"] = time.Now().Add(time.Hour * 30 * 164).Unix()

	rT, err := jwt.NewWithClaims(jwt.SigningMethodHS256, rClaims).SignedString([]byte(key))
	if err != nil {
		return "", "", err
	}

	return t, rT, nil
}
