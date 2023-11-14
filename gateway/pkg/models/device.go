package models

type RemoveDeviceReq struct {
	DeviceUUID string `json:"deviceUUID,omitempty" validate:"required"`
	UserUUID   string `json:"userUUID,omitempty" validate:"required"`
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
