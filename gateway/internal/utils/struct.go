package utils

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Decode[T any](r *http.Request) (*T, error) {
	var str T
	if err := json.NewDecoder(r.Body).Decode(&str); err != nil {
		return nil, err
	}

	v := validator.New()
	if err := v.Struct(str); err != nil {
		return nil, err
	}

	return &str, nil
}
