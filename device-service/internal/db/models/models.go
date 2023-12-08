package models

type Device struct {
	UUID         string
	Title        string
	Description  string
	Price        float32
	Manufacturer string
	Amount       uint32
}

type UpdateDevice struct {
	Title       string
	Description string
	Price       float32
	UUID        string
}

type GetByPrice struct {
	Min float32
	Max float32
}

type OrderDevice struct {
	DeviceUUID string
	Amount     uint32
}
