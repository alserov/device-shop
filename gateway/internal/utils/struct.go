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
		for _, v := range validator {
			if err := v(&str); err != nil {
				return nil, err
			}
		}
	}

	return &str, nil
}
