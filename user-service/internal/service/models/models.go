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
	Cash   float32
	TxUUID string
	Status int
}
