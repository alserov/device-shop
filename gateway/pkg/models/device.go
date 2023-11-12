package models

import "github.com/go-playground/validator/v10"

func Validate[T interface{}](s *T) error {
	v := validator.New()
	return v.Struct(s)
}

type Device struct {
	UUID         string `json:"uuid,omitempty" bson:"uuid"`
	Title        string `json:"title,omitempty" bson:"title" validate:"required,min=3"`
	Description  string `json:"description,omitempty" bson:"description"`
	Price        int32  `json:"price,omitempty" bson:"price" validate:"required,gt=0"`
	Manufacturer string `json:"manufacturer,omitempty" bson:"manufacturer" validate:"required"`
	Amount       uint   `bson:"amount,omitempty" bson:"amount" validate:"gt=0"`
}

type DeleteReq struct {
	UUID string `json:"uuid,omitempty" validate:"required"`
}

type RemoveReq struct {
	DeviceUUID string `json:"deviceUUID,omitempty" validate:"required"`
	UserUUID   string `json:"userUUID,omitempty" validate:"required"`
}

type UpdateReq struct {
	UUID        string `json:"uuid,omitempty" validate:"required"`
	Title       string `json:"title,omitempty" validate:"required,min=3"`
	Description string `json:"description,omitempty"`
	Price       int32  `json:"price,omitempty" validate:"required,gt=0"`
}

type GetAllReq struct {
	Index  *int32 `json:"index,omitempty" validate:"required,gt=-1"`
	Amount *int32 `json:"amount,omitempty" validate:"required,gt=0"`
}
