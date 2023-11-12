package entity

type Device struct {
	UUID         string  `json:"uuid,omitempty"`
	Title        string  `json:"title,omitempty" validate:"required,min=3"`
	Description  string  `json:"description,omitempty"`
	Price        float32 `json:"price,omitempty" validate:"required,gt=0"`
	Manufacturer string  `json:"manufacturer,omitempty" validate:"required"`
}

type UpdateDeviceReq struct {
	UUID        string
	Title       string
	Description string
	Price       float32
}
