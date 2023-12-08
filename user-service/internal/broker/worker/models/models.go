package models

type Request struct {
	TxUUID     string  `json:"txUUID"`
	OrderPrice float32 `json:"orderPrice"`
	UserUUID   string  `json:"userUUID"`
	Status     uint32  `json:"status"`
}

type Response struct {
	// 0 - failed
	// 1 - success
	Status  uint32 `json:"status"`
	Message string `json:"message"`
	Uuid    string `json:"uuid"`
}
