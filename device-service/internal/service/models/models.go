package models

type CreateDeviceReq struct {
	Title        string
	Description  string
	Price        float32
	Manufacturer string
	Amount       uint32
	UUID         string
}

type UpdateDeviceReq struct {
	Title       string
	Description string
	Price       float32
	UUID        string
}

type Device struct {
	UUID         string
	Title        string
	Description  string
	Price        float32
	Manufacturer string
	Amount       uint32
}

type GetAllDevicesReq struct {
	Amount uint32
	Index  uint32
}

type GetByPrice struct {
	Min float32
	Max float32
}

type IncreaseDeviceAmountReq struct {
	DeviceUUID string
	Amount     uint32
}
