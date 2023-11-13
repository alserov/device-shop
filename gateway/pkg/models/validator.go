package models

import "github.com/go-playground/validator/v10"

func Validate[T interface{}](s *T) error {
	v := validator.New()
	return v.Struct(s)
}
