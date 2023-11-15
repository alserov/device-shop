package entity

type Device struct {
	UUID         string  `json:"uuid,omitempty" bson:"uuid"`
	Title        string  `json:"title,omitempty" bson:"title" validate:"required,min=3"`
	Description  string  `json:"description,omitempty" bson:"description"`
	Price        float32 `json:"price,omitempty" bson:"price" validate:"required,gt=0"`
	Manufacturer string  `json:"manufacturer,omitempty" bson:"manufacturer" validate:"required"`
	Amount       uint32  `bson:"amount,omitempty" bson:"amount" validate:"gt=0"`
}

type UpdateDeviceReq struct {
	UUID        string  `json:"uuid,omitempty" validate:"required"`
	Title       string  `json:"title,omitempty" validate:"required,min=3"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price,omitempty" validate:"required,gt=0"`
}

type GetAllDevicesReq struct {
	Index  *int32 `json:"index,omitempty" validate:"required,gt=-1"`
	Amount *int32 `json:"amount,omitempty" validate:"required,gt=0"`
}

type RemoveDeviceReq struct {
	DeviceUUID string `json:"deviceUUID,omitempty" validate:"required"`
	UserUUID   string `json:"userUUID,omitempty" validate:"required"`
}
