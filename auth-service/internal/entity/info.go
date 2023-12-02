package entity

type GetUserInfoReq struct {
	UUID string
}

type CheckIfAdminReq struct {
	UUID string
}

type GetUserInfoRes struct {
	Username string
	Email    string
	UUID     string
	Cash     float32
}

type CheckIfAdminRes struct {
	IsAdmin bool
}
