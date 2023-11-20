package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alserov/device-shop/auth-service/internal/entity"
	pb "github.com/alserov/device-shop/proto/gen"
	"google.golang.org/grpc/status"
	"net/http"
)

type Repository interface {
	Signup(ctx context.Context, req *pb.SignupReq, info *entity.SignupAdditional) error
	Login(ctx context.Context, req *pb.LoginReq, rToken string) (string, error)
	GetInfoByUsername(ctx context.Context, uname string) (string, string, error)
}

type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) Signup(ctx context.Context, req *pb.SignupReq, info *entity.SignupAdditional) error {
	query := `INSERT INTO users (uuid,username,password,email,cash,refresh_token,role,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`

	_, err := r.db.Exec(query, info.UUID, req.Username, req.Password, req.Email, info.Cash, info.RefreshToken, info.Role, info.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Login(ctx context.Context, req *pb.LoginReq, rToken string) (string, error) {
	query := `UPDATE users SET refresh_token = $1 WHERE username = $2 RETURNING uuid`

	var uuid string
	if err := r.db.QueryRow(query, rToken, req.Username).Scan(&uuid); err != nil {
		return "", err
	}

	return uuid, nil
}

func (r *repo) GetInfoByUsername(ctx context.Context, uname string) (string, string, error) {
	query := `SELECT password, role FROM users WHERE username = $1 LIMIT 1`

	var (
		password string
		role     string
	)
	if err := r.db.QueryRow(query, uname).Scan(&password, &role); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", status.Error(http.StatusBadRequest, "user not found")
		}
		return "", "", err
	}
	return password, role, nil
}
