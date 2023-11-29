package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/auth-service/internal/broker"
	"github.com/alserov/device-shop/auth-service/internal/db"
	"github.com/alserov/device-shop/auth-service/internal/db/postgres"
	"github.com/alserov/device-shop/auth-service/internal/utils"
	conv "github.com/alserov/device-shop/auth-service/internal/utils/proto_converter"
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type service struct {
	auth        db.AuthRepo
	emailTopic  string
	emailBroker string
}

func New(pg *sql.DB, topic string, brokerAddr string) pb.AuthServer {
	return &service{
		auth:        postgres.NewAuthRepo(pg),
		emailTopic:  topic,
		emailBroker: brokerAddr,
	}
}

const (
	defaultRole   = "user"
	kafkaClientID = "SIGNUP_RPC"
)

func (s *service) Signup(ctx context.Context, req *pb.SignupReq) (*pb.SignupRes, error) {
	if _, _, err := s.auth.GetPasswordAndRoleByUsername(ctx, req.Username); err == nil {
		return &pb.SignupRes{}, status.Error(codes.AlreadyExists, "user with this username already exists")
	}

	now := time.Now().UTC() /*createdAt*/

	token, rToken, err := utils.GenerateTokens(defaultRole)
	if err != nil {
		return &pb.SignupRes{}, err
	}

	r := conv.SignupReqToRepoStruct(req)
	r.Password, err = utils.HashPassword(req.Password)
	if err != nil {
		return &pb.SignupRes{}, err
	}

	info := db.SignupInfo{
		UUID:         uuid.New().String(),
		Cash:         0,
		Role:         defaultRole,
		CreatedAt:    &now,
		RefreshToken: rToken,
	}

	if err = s.auth.Signup(ctx, r, info); err != nil {
		return &pb.SignupRes{}, err
	}

	producer, err := broker.NewProducer([]string{s.emailBroker}, kafkaClientID)
	if err != nil {
		return &pb.SignupRes{}, fmt.Errorf("failed to send a message to: %s", req.Email)
	}

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Value: sarama.StringEncoder(req.Email),
		Topic: s.emailTopic,
	})
	if err != nil {
		return &pb.SignupRes{}, fmt.Errorf("failed to send a message to: %s", req.Email)
	}

	return &pb.SignupRes{
		Username:     req.Username,
		Email:        req.Email,
		UUID:         info.UUID,
		Cash:         info.Cash,
		RefreshToken: info.RefreshToken,
		Token:        token,
	}, nil
}

func (s *service) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	password, role, err := s.auth.GetPasswordAndRoleByUsername(ctx, req.Username)
	if err != nil {
		return &pb.LoginRes{}, err
	}

	if err = utils.CheckPassword(req.Password, password); err != nil {
		return nil, err
	}

	token, rToken, err := utils.GenerateTokens(role)
	if err != nil {
		return &pb.LoginRes{}, err
	}

	r := conv.LoginReqToRepoStruct(req)
	userUUID, err := s.auth.Login(ctx, r, rToken)
	if err != nil {
		return &pb.LoginRes{}, err
	}

	return &pb.LoginRes{
		RefreshToken: rToken,
		Token:        token,
		UUID:         userUUID,
	}, nil
}
