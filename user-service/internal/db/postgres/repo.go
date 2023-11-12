package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alserov/device-shop/user-service/internal/entity"
	"google.golang.org/grpc/status"
	"net/http"
)

type Repository interface {
	Signup(ctx context.Context, req *entity.User) error
	Login(ctx context.Context, req *entity.RepoLoginReq) error
	CheckIfExistsByUsername(ctx context.Context, req string) (bool, error)
	FindByUsername(ctx context.Context, req string) (*entity.User, error)
}

func NewRepo(db *sql.DB) Repository {
	return &repo{
		db: db,
	}
}

type repo struct {
	db *sql.DB
}

func (r *repo) Signup(ctx context.Context, req *entity.User) error {
	query := `INSERT INTO users (uuid,username,password,email,cash,refresh_token,role,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7, $8)`

	_, err := r.db.Exec(query, req.UUID, req.Username, req.Password, req.Email, req.Cash, req.RefreshToken, req.Role, req.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Login(ctx context.Context, req *entity.RepoLoginReq) error {
	query := `UPDATE users SET refresh_token = $1 WHERE username = $2`

	if _, err := r.db.Exec(query, req.RefreshToken, req.Username); err != nil {
		return err
	}

	return nil
}

func (r *repo) FindByUsername(ctx context.Context, uname string) (*entity.User, error) {
	query := `SELECT * FROM users WHERE username = $1`

	row := r.db.QueryRow(query, uname)

	var foundUser entity.User
	err := row.Scan(&foundUser.UUID, &foundUser.Username, &foundUser.Password, &foundUser.Email, &foundUser.Cash, &foundUser.RefreshToken, &foundUser.Role, &foundUser.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(http.StatusBadRequest, "user not found")
	}
	if err != nil {
		return nil, err
	}

	return &foundUser, nil
}

func (r *repo) CheckIfExistsByUsername(ctx context.Context, uname string) (bool, error) {
	query := `SELECT count(*) FROM users WHERE username = $1`

	row := r.db.QueryRow(query, uname)

	var c int
	if err := row.Scan(&c); err != nil {
		return false, err
	}
	if c != 0 {
		return true, nil
	}

	return false, nil
}
