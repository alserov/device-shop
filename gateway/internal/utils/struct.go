package utils

import (
	"encoding/json"
	"net/http"
)

func Decode[T any](r *http.Request, validator ...func(*T) error) (*T, error) {
	var str T
	if err := json.NewDecoder(r.Body).Decode(&str); err != nil {
		return nil, err
	}

	if len(validator) > 0 {
		if err := validator[0](&str); err != nil {
			return nil, err
		}
	}

	return &str, nil
}
