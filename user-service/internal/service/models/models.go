package models

type SignupReq struct {
	Username string
	Email    string
	Password string
}

type SignupRes struct {
	Username     string
	Email        string
	UUID         string
	Cash         float32
	RefreshToken string
	Token        string
}

type LoginReq struct {
	Username string
	Password string
}

type LoginRes struct {
	RefreshToken string
	Token        string
	UUID         string
}

type GetUserInfoRes struct {
	Username string
	Email    string
	UUID     string
	Cash     float32
}

type BalanceReq struct {
	Cash     float32
	UserUUID string
}

type WorkerBalanceReq struct {
	TxUUID     string  `json:"txUUID"`
	OrderPrice float32 `json:"orderPrice"`
	UserUUID   string  `json:"userUUID"`
	Status     uint32  `json:"status"`
}

type TxResponse struct {
	// 0 - failed
	// 1 - success
	Status  uint32 `json:"status"`
	Message string `json:"message"`
	Uuid    string `json:"uuid"`
}
