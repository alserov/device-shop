package validation

import (
	"github.com/alserov/device-shop/proto/gen/user"
	"github.com/stretchr/testify/require"
	"testing"
)

type test[T any] struct {
	Struct      T
	ExpectError bool
}

func TestValidateAuth(t *testing.T) {
	testSignupReq(t)
	testLoginReq(t)
}

func testLoginReq(t *testing.T) {
	testSignup := []test[user.LoginReq]{
		{
			Struct: user.LoginReq{
				Username: "username",
				Password: "",
			},
			ExpectError: true,
		},
		{
			Struct: user.LoginReq{
				Username: "",
				Password: "password",
			},
			ExpectError: true,
		},
		{
			Struct: user.LoginReq{
				Username: "username",
				Password: "password",
			},
			ExpectError: false,
		},
	}

	a := &auth{}

	for _, tc := range testSignup {
		err := a.ValidateLoginReq(&tc.Struct)
		switch tc.ExpectError {
		case true:
			require.Error(t, err)
		default:
			require.NoError(t, err)
		}
	}
}

func testSignupReq(t *testing.T) {
	testSignup := []test[user.SignupReq]{
		{
			Struct: user.SignupReq{
				Email:    "email",
				Username: "username",
				Password: "",
			},
			ExpectError: true,
		},
		{
			Struct: user.SignupReq{
				Email:    "email",
				Username: "",
				Password: "password",
			},
			ExpectError: true,
		},
		{
			Struct: user.SignupReq{
				Email:    "",
				Username: "username",
				Password: "password",
			},
			ExpectError: true,
		},
		{
			Struct: user.SignupReq{
				Email:    "email",
				Username: "username",
				Password: "password",
			},
			ExpectError: false,
		},
	}

	a := &auth{}

	for _, tc := range testSignup {
		err := a.ValidateSignupReq(&tc.Struct)
		switch tc.ExpectError {
		case true:
			require.Error(t, err)
		default:
			require.NoError(t, err)
		}
	}
}

func TestValidateBalance(t *testing.T) {
	tests := []test[user.BalanceReq]{
		{
			Struct: user.BalanceReq{
				UserUUID: "",
				Cash:     1,
			},
			ExpectError: true,
		},
		{
			Struct: user.BalanceReq{
				UserUUID: "uuid",
				Cash:     0,
			},
			ExpectError: true,
		},
		{
			Struct: user.BalanceReq{
				UserUUID: "uuid",
				Cash:     1,
			},
			ExpectError: false,
		},
	}
	b := &balance{}

	for _, tc := range tests {
		err := b.ValidateBalanceReq(&tc.Struct)
		switch tc.ExpectError {
		case true:
			require.Error(t, err)
		default:
			require.NoError(t, err)
		}
	}
}
