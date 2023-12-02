package service

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

type GetDeviceByTitleReq struct {
	Title string
}

type GetDeviceByUUIDReq struct {
	UUID string
}

type GetByManufacturer struct {
	Manufacturer string
}

type GetByPrice struct {
	Min float32
	Max float32
}
