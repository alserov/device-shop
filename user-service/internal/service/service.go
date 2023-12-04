package service

import (
	"context"
	"database/sql"
	"github.com/IBM/sarama"

	"github.com/alserov/device-shop/user-service/internal/broker"
	"github.com/alserov/device-shop/user-service/internal/db"
	repo "github.com/alserov/device-shop/user-service/internal/db/models"
	"github.com/alserov/device-shop/user-service/internal/db/postgres"
	"github.com/alserov/device-shop/user-service/internal/service/models"
	"github.com/alserov/device-shop/user-service/internal/utils"
	"github.com/alserov/device-shop/user-service/internal/utils/converter"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"time"
)

type service struct {
	log *slog.Logger
	db  db.UserRepo

	emailBroker string
	emailTopic  string

	conv *converter.ServiceConverter
}

func NewService(pg *sql.DB, log *slog.Logger, emailBroker string, emailTopic string) Service {
	return &service{
		log:         log,
		db:          postgres.NewRepo(pg, log),
		conv:        converter.NewServiceConverter(),
		emailBroker: emailBroker,
		emailTopic:  emailTopic,
	}
}

type Service interface {
	TopUpBalance(ctx context.Context, req models.BalanceReq) (float32, error)
	DebitBalance(ctx context.Context, req models.BalanceReq) (float32, error)

	Signup(ctx context.Context, req models.SignupReq) (models.SignupRes, error)
	Login(ctx context.Context, req models.LoginReq) (models.LoginRes, error)

	GetUserInfo(ctx context.Context, uuid string) (models.GetUserInfoRes, error)
	CheckIfAdmin(ctx context.Context, uuid string) (bool, error)
}

func (s *service) TopUpBalance(ctx context.Context, req models.BalanceReq) (float32, error) {
	cash, err := s.db.TopUpBalance(ctx, repo.BalanceReq{
		Cash:     req.Cash,
		UserUUID: req.UserUUID,
	})
	if err != nil {
		return 0, err
	}

	return cash, nil
}

func (s *service) DebitBalance(ctx context.Context, req models.BalanceReq) (float32, error) {
	cash, err := s.db.DebitBalance(ctx, repo.BalanceReq{
		Cash:     req.Cash,
		UserUUID: req.UserUUID,
	})
	if err != nil {
		return 0, err
	}

	return cash, nil
}

const (
	defaultRole   = "user"
	kafkaClientID = "SIGNUP_RPC"
	internalErr   = "internal error"
)

func (s *service) Signup(ctx context.Context, req models.SignupReq) (models.SignupRes, error) {
	op := "service.Signup"
	// err == nil => means that user already exists
	if _, _, err := s.db.GetPasswordAndRoleByUsername(ctx, req.Username); err == nil {
		return models.SignupRes{}, status.Error(codes.AlreadyExists, "user with this username already exists")
	}

	now := time.Now().UTC() /*createdAt*/

	token, rToken, err := utils.GenerateTokens(defaultRole)
	if err != nil {
		s.log.Error("failed to generate tokens", slog.String("error", err.Error()), slog.String("op", op))
		return models.SignupRes{}, status.Error(codes.Internal, internalErr)
	}

	req.Password, err = utils.HashPassword(req.Password)
	if err != nil {
		s.log.Error("failed to hash password", slog.String("error", err.Error()), slog.String("op", op))
		return models.SignupRes{}, status.Error(codes.Internal, internalErr)
	}

	info := repo.SignupInfo{
		UUID:         uuid.New().String(),
		Cash:         0,
		Role:         defaultRole,
		CreatedAt:    &now,
		RefreshToken: rToken,
	}
	if err = s.db.Signup(ctx, s.conv.Auth.SignupReqToRepo(req), info); err != nil {
		return models.SignupRes{}, err
	}

	producer, err := broker.NewProducer([]string{s.emailBroker}, kafkaClientID)
	if err != nil {
		s.log.Error("failed to initialize new kafka producer", slog.String("error", err.Error()), slog.String("op", op))
	}
	if err == nil {
		_, _, err = producer.SendMessage(&sarama.ProducerMessage{
			Value: sarama.StringEncoder(req.Email),
			Topic: s.emailTopic,
		})
		if err != nil {
			s.log.Error("failed to send message", slog.String("error", err.Error()), slog.String("op", op))
		}
	}

	return s.conv.Auth.SignupResToService(req, info, token), nil
}

func (s *service) Login(ctx context.Context, req models.LoginReq) (models.LoginRes, error) {
	op := "service.Login"

	password, role, err := s.db.GetPasswordAndRoleByUsername(ctx, req.Username)
	if err != nil {
		return models.LoginRes{}, err
	}

	if err = utils.CheckPassword(req.Password, password); err != nil {
		return models.LoginRes{}, err
	}

	token, rToken, err := utils.GenerateTokens(role)
	if err != nil {
		s.log.Error("failed to generate tokens", slog.String("error", err.Error()), slog.String("op", op))
		return models.LoginRes{}, status.Error(codes.Internal, internalErr)
	}

	userUUID, err := s.db.Login(ctx, s.conv.Auth.LoginReqToRepo(req), rToken)
	if err != nil {
		return models.LoginRes{}, err
	}

	return s.conv.Auth.LoginResToService(rToken, token, userUUID), nil
}

func (s *service) GetUserInfo(ctx context.Context, uuid string) (models.GetUserInfoRes, error) {
	info, err := s.db.GetUserInfo(ctx, uuid)
	if err != nil {
		return models.GetUserInfoRes{}, err
	}

	return s.conv.Info.UserInfoResToService(info), nil
}

func (s *service) CheckIfAdmin(ctx context.Context, uuid string) (bool, error) {
	isAdmin, err := s.db.CheckIfAdmin(ctx, uuid)
	if err != nil {
		return false, err
	}

	return isAdmin, nil
}
