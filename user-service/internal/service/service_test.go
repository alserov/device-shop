package service

import (
	"context"
	"errors"
	"github.com/alserov/device-shop/user-service/internal/utils"

	mailmock "github.com/alserov/device-shop/user-service/internal/broker/mail/mocks"
	repomock "github.com/alserov/device-shop/user-service/internal/db/mocks"
	"github.com/alserov/device-shop/user-service/internal/service/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTopUpBalance(t *testing.T) {
	balanceReq := models.BalanceReq{
		UserUUID: "uuid",
		Cash:     10,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := repomock.NewMockUserRepo(ctrl)
	store.EXPECT().TopUpBalance(gomock.Any(), gomock.Any()).Return(float32(10), nil).Times(1)

	s := NewService(store, nil, nil, nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	balance, err := s.TopUpBalance(ctx, balanceReq)
	require.NoError(t, err)
	require.Equal(t, balanceReq.Cash, balance)
}

func TestDebitBalance(t *testing.T) {
	balanceReq := models.BalanceReq{
		UserUUID: "uuid",
		Cash:     10,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := repomock.NewMockUserRepo(ctrl)
	store.EXPECT().DebitBalance(gomock.Any(), gomock.Any()).Return(float32(0), nil).Times(1)

	s := NewService(store, nil, nil, nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	balance, err := s.DebitBalance(ctx, balanceReq)
	require.NoError(t, err)
	require.Equal(t, balanceReq.Cash, balance+10)
}

func TestSignup(t *testing.T) {
	signupReq := models.SignupReq{
		Email:    "email",
		Username: "username",
		Password: "password",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := repomock.NewMockUserRepo(ctrl)
	store.EXPECT().Signup(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
	store.EXPECT().GetPasswordAndRoleByUsername(gomock.Any(), gomock.Any()).Return("", "", errors.New("user not found")).Times(1)

	mail := mailmock.NewMockEmailer(ctrl)
	mail.EXPECT().Send(gomock.Any()).Return(nil).Times(1)

	s := NewService(store, nil, mail, nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	res, err := s.Signup(ctx, signupReq)
	require.NoError(t, err)

	require.Equal(t, res.Username, signupReq.Username)
	require.Equal(t, res.Email, signupReq.Email)
	require.NotEmpty(t, res.UUID)
	require.Zero(t, res.Cash)
	require.NotEmpty(t, res.Token)
	require.NotEmpty(t, res.RefreshToken)
}

func TestLogin(t *testing.T) {
	loginReq := models.LoginReq{
		Username: "username",
		Password: "password",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	hashedPassword, err := utils.HashPassword(loginReq.Password)
	require.NoError(t, err)

	store := repomock.NewMockUserRepo(ctrl)
	store.EXPECT().Login(gomock.Any(), gomock.Any(), gomock.Any()).Return("uuid", nil).Times(1)
	store.EXPECT().GetPasswordAndRoleByUsername(gomock.Any(), gomock.Any()).Return(hashedPassword, "user", nil).Times(1)

	s := NewService(store, nil, nil, nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	uuid, err := s.Login(ctx, loginReq)
	require.NoError(t, err)

	require.NotEmpty(t, uuid)
}
