package proto_converter

import (
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/stretchr/testify/assert"
	"testing"
)

type test[T any] struct {
	in T
}

func TestSignupReqToRepoStruct(t *testing.T) {
	tests := []test[pb.SignupReq]{
		{
			in: pb.SignupReq{
				Username: "username",
				Email:    "email",
				Password: "password",
			},
		},
	}

	for _, tc := range tests {
		convertedStruct := SignupReqToRepoStruct(&tc.in)

		assert.Equal(t, tc.in.Username, convertedStruct.Username)
		assert.Equal(t, tc.in.Email, convertedStruct.Email)
		assert.Equal(t, tc.in.Password, convertedStruct.Password)
	}
}

func TestLoginReqToRepoStruct(t *testing.T) {
	tests := []test[pb.LoginReq]{
		{
			in: pb.LoginReq{
				Username: "username",
				Password: "password",
			},
		},
	}

	for _, tc := range tests {
		convertedStruct := LoginReqToRepoStruct(&tc.in)

		assert.Equal(t, tc.in.Username, convertedStruct.Username)
		assert.Equal(t, tc.in.Password, convertedStruct.Password)
	}
}
