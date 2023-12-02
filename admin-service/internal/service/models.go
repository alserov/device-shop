package service

type CreateDeviceReq struct {
	Title        string
	Description  string
	Price        float32
	Manufacturer string
	Amount       uint32
}

type DeleteDeviceReq struct {
	UUID string
}

type UpdateDeviceReq struct {
	Title       string
	Description string
	Price       float32
	UUID        string
}
