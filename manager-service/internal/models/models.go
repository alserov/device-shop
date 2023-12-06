package models

type Tx struct {
	Uuid   string `json:"uuid"`
	AllTxs uint32 `json:"allTxs"`
}

type TxRequest struct {
	UserUUID     string          `json:"userUUID"`
	OrderPrice   float32         `json:"orderPrice"`
	OrderDevices []*OrderDevices `json:"orderDevices"`
}

type OrderDevices struct {
	DeviceUUID string `json:"deviceUUID"`
	Amount     uint32 `json:"amount"`
}

type TxResponse struct {
	// 0 - failed
	// 1 - success
	Status  uint32 `json:"status"`
	Message string `json:"message"`
	Uuid    string `json:"uuid"`
}
