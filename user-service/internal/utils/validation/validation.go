package validation

import (
	"github.com/alserov/device-shop/proto/gen/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	emptyUUID      = "uuid can not be empty"
	emptyPassword  = "password can not be empty"
	emptyUsername  = "username can not be empty"
	invalidBalance = "balance can not be less or equal to 0"
)

type Validator struct {
	Auth    Auth
	Info    Info
	Balance Balance
}

type auth struct{}
type Auth interface {
	ValidateSignupReq(r *user.SignupReq) error
	ValidateLoginReq(r *user.LoginReq) error
}

type info struct{}
type Info interface {
	ValidateGetUserInfoReq(r *user.GetUserInfoReq) error
	ValidateCheckIfAdminReq(r *user.CheckIfAdminReq) error
}

type balance struct{}
type Balance interface {
	ValidateBalanceReq(r *user.BalanceReq) error
}

func NewValidator() *Validator {
	return &Validator{
		Auth:    &auth{},
		Info:    &info{},
		Balance: &balance{},
	}
}

func (b *balance) ValidateBalanceReq(r *user.BalanceReq) error {
	if r.GetUserUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyUUID)
	}

	if r.GetCash() <= 0 {
		return status.Error(codes.InvalidArgument, invalidBalance)
	}

	return nil
}

func (a *auth) ValidateSignupReq(r *user.SignupReq) error {
	if r.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, emptyPassword)
	}

	if r.GetUsername() == "" {
		return status.Error(codes.InvalidArgument, emptyUsername)
	}

	if r.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email can not be empty")
	}

	return nil
}

func (a *auth) ValidateLoginReq(r *user.LoginReq) error {
	if r.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, emptyPassword)
	}

	if r.GetUsername() == "" {
		return status.Error(codes.InvalidArgument, emptyUsername)
	}

	return nil
}

func (i *info) ValidateGetUserInfoReq(r *user.GetUserInfoReq) error {
	if r.GetUserUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyUUID)
	}

	return nil
}

func (i *info) ValidateCheckIfAdminReq(r *user.CheckIfAdminReq) error {
	if r.GetUserUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyUUID)
	}

	return nil
}
