package service

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/auth-service/internal/db"
	"github.com/alserov/device-shop/auth-service/internal/db/postgres"
	"github.com/alserov/device-shop/auth-service/internal/entity"
	"github.com/alserov/device-shop/auth-service/internal/utils"
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

type service struct {
	auth db.AuthRepo
}

func New(pg *sql.DB) pb.AuthServer {
	return &service{
		auth: postgres.NewAuthRepo(pg),
	}
}

func (s *service) Signup(ctx context.Context, req *pb.SignupReq) (*pb.SignupRes, error) {
	if _, _, err := s.auth.GetPasswordAndRoleByUsername(ctx, req.Username); err == nil {
		return &pb.SignupRes{}, err
	}

	now := time.Now().UTC()
	token, rToken, err := utils.GenerateTokens("user")
	if err != nil {
		return &pb.SignupRes{}, err
	}

	info := &entity.SignupAdditional{
		UUID:         uuid.New().String(),
		Cash:         0,
		Role:         "user",
		CreatedAt:    &now,
		RefreshToken: rToken,
	}

	req.Password, err = utils.HashPassword(req.Password)
	if err != nil {
		return &pb.SignupRes{}, err
	}

	if err = s.auth.Signup(ctx, req, info); err != nil {
		return &pb.SignupRes{}, err
	}

	//if err = utils.SendEmail(r.Email); err != nil {
	//	log.Println("FAILED TO SEND EMAIL: ", err.Error())
	//}

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
		return nil, status.Error(http.StatusBadRequest, "invalid password")
	}

	token, rToken, err := utils.GenerateTokens(role)
	if err != nil {
		return &pb.LoginRes{}, err
	}

	userUUID, err := s.auth.Login(ctx, req, rToken)
	if err != nil {
		return &pb.LoginRes{}, err
	}

	return &pb.LoginRes{
		RefreshToken: rToken,
		Token:        token,
		UUID:         userUUID,
	}, nil
}
