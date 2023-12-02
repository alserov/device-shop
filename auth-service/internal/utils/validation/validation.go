package validation

import (
	"github.com/alserov/device-shop/proto/gen/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	emptyUUID     = "uuid can not be empty"
	emptyPassword = "password can not be empty"
	emptyUsername = "username can not be empty"
)

func ValidateSignupReq(r *auth.SignupReq) error {
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

func ValidateLoginReq(r *auth.LoginReq) error {
	if r.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, emptyPassword)
	}

	if r.GetUsername() == "" {
		return status.Error(codes.InvalidArgument, emptyUsername)
	}

	return nil
}

func ValidateGetUserInfoReq(r *auth.GetUserInfoReq) error {
	if r.GetUserUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyUUID)
	}

	return nil
}

func ValidateCheckIfAdminReq(r *auth.CheckIfAdminReq) error {
	if r.GetUserUUID() == "" {
		return status.Error(codes.InvalidArgument, emptyUUID)
	}

	return nil
}
