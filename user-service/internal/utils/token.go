package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return "", "", status.Error(codes.Internal, err.Error())
	}

	rClaims := jwt.MapClaims{}
	rClaims["exp"] = time.Now().Add(time.Hour * 30 * 164).Unix()

	rT, err := jwt.NewWithClaims(jwt.SigningMethodHS256, rClaims).SignedString([]byte(key))
	if err != nil {
		return "", "", status.Error(codes.Internal, err.Error())
	}

	return t, rT, nil
}
