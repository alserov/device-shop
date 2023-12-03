package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/auth-service/internal/db"
	"github.com/alserov/device-shop/auth-service/internal/db/postgres"
	"github.com/alserov/device-shop/proto/gen/auth"
	"github.com/alserov/device-shop/user-service/internal/broker"
	repo "github.com/alserov/device-shop/user-service/internal/db/models"
	"github.com/alserov/device-shop/user-service/internal/service/models"
	utils2 "github.com/alserov/device-shop/user-service/internal/utils"
	"github.com/alserov/device-shop/user-service/internal/utils/converter"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Service interface {
	Signup(ctx context.Context, req *models.SignupReq) (*models.SignupRes, error)
	Login(ctx context.Context, req *models.LoginReq) (*models.LoginRes, error)
	GetUserInfo(ctx context.Context, uuid string) (*models.GetUserInfoRes, error)
	CheckIfAdmin(ctx context.Context, uuid string) (bool, error)
}

type service struct {
	auth.UnimplementedAuthServer
	db          db.AuthRepo
	emailTopic  string
	emailBroker string
}

func NewService(pg *sql.DB, topic string, brokerAddr string) Service {
	return &service{
		db:          postgres.NewAuthRepo(pg),
		emailTopic:  topic,
		emailBroker: brokerAddr,
	}
}

const (
	defaultRole   = "user"
	kafkaClientID = "SIGNUP_RPC"
)

func (s *service) Signup(ctx context.Context, req *models.SignupReq) (*models.SignupRes, error) {
	if _, _, err := s.db.GetPasswordAndRoleByUsername(ctx, req.Username); err == nil {
		return nil, status.Error(codes.AlreadyExists, "user with this username already exists")
	}

	now := time.Now().UTC() /*createdAt*/

	token, rToken, err := utils2.GenerateTokens(defaultRole)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to generate tokens: %v", err))
	}

	req.Password, err = utils2.HashPassword(req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to hash password: %v", err))
	}

	info := repo.SignupInfo{
		UUID:         uuid.New().String(),
		Cash:         0,
		Role:         defaultRole,
		CreatedAt:    &now,
		RefreshToken: rToken,
	}
	if err = s.db.Signup(ctx, converter.ServiceSignupReqToRepo(req), info); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to insert user: %v", err))
	}

	producer, err := broker.NewProducer([]string{s.emailBroker}, kafkaClientID)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to send a message to: %s", req.Email))
	}

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Value: sarama.StringEncoder(req.Email),
		Topic: s.emailTopic,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to send a message to: %s", req.Email))
	}

	return &models.SignupRes{
		Username:     req.Username,
		Email:        req.Email,
		UUID:         info.UUID,
		Cash:         info.Cash,
		RefreshToken: info.RefreshToken,
		Token:        token,
	}, nil
}

func (s *service) Login(ctx context.Context, req *models.LoginReq) (*models.LoginRes, error) {
	password, role, err := s.db.GetPasswordAndRoleByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	if err = utils2.CheckPassword(req.Password, password); err != nil {
		return nil, err
	}

	token, rToken, err := utils2.GenerateTokens(role)
	if err != nil {
		return nil, err
	}
	userUUID, err := s.db.Login(ctx, converter.ServiceLoginReqToRepo(req), rToken)
	if err != nil {
		return nil, err
	}

	return &models.LoginRes{
		RefreshToken: rToken,
		Token:        token,
		UUID:         userUUID,
	}, nil
}

func (s *service) GetUserInfo(ctx context.Context, uuid string) (*models.GetUserInfoRes, error) {
	info, err := s.db.GetUserInfo(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return converter.RepoUserInfoResToService(info), nil
}

func (s *service) CheckIfAdmin(ctx context.Context, uuid string) (bool, error) {
	isAdmin, err := s.db.CheckIfAdmin(ctx, uuid)
	if err != nil {
		return false, err
	}

	return isAdmin, nil
}
