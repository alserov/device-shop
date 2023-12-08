package models

type Response struct {
	// -1 - failed (server error)
	// 0 - failed (user error)
	// 1 - success
	Status  uint32 `json:"status"`
	Message string `json:"message"`
	Uuid    string `json:"uuid"`
}

type Request struct {
	OrderDevices []*OrderDevice
	TxUUID       string
	Status       uint32
}

type OrderDevice struct {
	DeviceUUID string
	Amount     uint32
}
