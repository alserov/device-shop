package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/alserov/device-shop/user-service/internal/db"
	"github.com/alserov/device-shop/user-service/internal/db/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewRepo(db *sql.DB) db.UserRepo {
	return &repo{
		db: db,
	}
}

type repo struct {
	db *sql.DB
}

func (r *repo) GetInfo(_ context.Context, userUUID string) (models.GetUserInfoRes, error) {
	query := `SELECT username,email,uuid,cash FROM users WHERE uuid = $1`

	var info models.GetUserInfoRes
	if err := r.db.QueryRow(query, userUUID).Scan(&info.Username, &info.Email, &info.UUID, &info.Cash); err != nil {
		return models.GetUserInfoRes{}, err
	}

	return info, nil
}

func (r *repo) TopUpBalance(_ context.Context, req models.BalanceReq) (float32, error) {
	query := `UPDATE users SET cash = cash + $1 WHERE uuid = $2 RETURNING cash`

	var cash float32
	if err := r.db.QueryRow(query, req.Cash, req.UserUUID).Scan(&cash); err != nil {
		return 0, err
	}

	return cash, nil
}

func (r *repo) DebitBalance(_ context.Context, req models.BalanceReq) (float32, error) {
	query := `UPDATE users SET cash = cash - $1 WHERE uuid = $2 RETURNING cash`

	var cash float32
	if err := r.db.QueryRow(query, req.Cash, req.UserUUID).Scan(&cash); err != nil {
		return 0, err
	}

	return cash, nil
}

func (r *repo) Signup(_ context.Context, req *models.SignupReq, info models.SignupInfo) error {
	query := `INSERT INTO users (uuid,username,password,email,cash,refresh_token,role,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`

	_, err := r.db.Exec(query, info.UUID, req.Username, req.Password, req.Email, info.Cash, info.RefreshToken, info.Role, info.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Login(_ context.Context, req *models.LoginReq, rToken string) (string, error) {
	query := `UPDATE users SET refresh_token = $1 WHERE username = $2 RETURNING uuid`

	var uuid string
	if err := r.db.QueryRow(query, rToken, req.Username).Scan(&uuid); err != nil {
		return "", status.Error(codes.Internal, err.Error())
	}

	return uuid, nil
}

func (r *repo) GetPasswordAndRoleByUsername(_ context.Context, uname string) (string, string, error) {
	query := `SELECT password, role FROM users WHERE username = $1 LIMIT 1`

	var (
		password string
		role     string
	)
	if err := r.db.QueryRow(query, uname).Scan(&password, &role); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", status.Error(codes.NotFound, fmt.Sprintf("user: %s not found", uname))
		}
		return "", "", status.Error(codes.Internal, fmt.Sprintf("failed to GetPasswordAndRoleByUsername: %v", err))
	}
	return password, role, nil
}

func (r *repo) GetUserInfo(ctx context.Context, uuid string) (*models.GetUserInfoRes, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repo) CheckIfAdmin(ctx context.Context, uuid string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
