package models

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/go-playground/validator/v10"
)

type SignupReq struct {
	Username string `json:"username,omitempty" validate:"min=3,max=50,required"`
	Password string `json:"password,omitempty" validate:"min=5,max=100,required"`
	Email    string `json:"email,omitempty" validate:"required"`
}

func (sq *SignupReq) Validate() error {
	v := validator.New()

	if err := v.Struct(sq); err != nil {
		return err
	}

	if valid := govalidator.IsEmail(sq.Email); !valid {
		return errors.New("invalid email")
	}

	return nil
}

type SignupRes struct {
	UUID     string `json:"UUID,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Cash     int    `json:"cash,omitempty"`
}

type LoginReq struct {
	Username string `json:"username,omitempty" validate:"min=3,max=50,required"`
	Password string `json:"password,omitempty" validate:"min=5,max=100,required"`
}
