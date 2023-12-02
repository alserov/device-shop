package entity

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
