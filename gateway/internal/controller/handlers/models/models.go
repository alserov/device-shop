package models

type Service struct {
	Addr string
}

type Services struct {
	Device Service
	User   Service
	Auth   Service
	Order  Service
}
