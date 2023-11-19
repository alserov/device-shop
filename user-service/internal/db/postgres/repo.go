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
	Signup(context.Context, *entity.User) error
	Login(context.Context, *entity.RepoLoginReq) (*entity.RepoLoginRes, error)
	GetInfo(context.Context, string) (*entity.RepoGetInfoRes, error)
	TopUpBalance(context.Context, *entity.TopUpBalanceReq) (float32, error)
	DebitBalance(context.Context, *entity.DebitBalanceReq) (float32, error)
	CheckIfExistsByUsername(context.Context, string) (bool, error)
	FindByUsername(context.Context, string) (*entity.User, error)
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

func (r *repo) Login(ctx context.Context, req *entity.RepoLoginReq) (*entity.RepoLoginRes, error) {
	query := `UPDATE users SET refresh_token = $1 WHERE username = $2 RETURNING uuid`

	var (
		uuid string
	)
	if err := r.db.QueryRow(query, req.RefreshToken, req.Username).Scan(&uuid); err != nil {
		return &entity.RepoLoginRes{}, err
	}

	return &entity.RepoLoginRes{
		UUID: uuid,
	}, nil
}

func (r *repo) GetInfo(ctx context.Context, userUUID string) (*entity.RepoGetInfoRes, error) {
	query := `SELECT username,email,uuid,cash FROM users WHERE uuid = $1`

	var info entity.RepoGetInfoRes
	if err := r.db.QueryRow(query, userUUID).Scan(&info.Username, &info.Email, &info.UUID, &info.Cash); err != nil {
		return nil, err
	}

	return &info, nil
}

func (r *repo) TopUpBalance(ctx context.Context, req *entity.TopUpBalanceReq) (float32, error) {
	query := `UPDATE users SET cash = cash + $1 WHERE uuid = $2 RETURNING cash`

	var cash float32
	if err := r.db.QueryRow(query, req.Cash, req.UserUUID).Scan(&cash); err != nil {
		return 0, err
	}

	return cash, nil
}

func (r *repo) DebitBalance(ctx context.Context, req *entity.DebitBalanceReq) (float32, error) {
	query := `UPDATE users SET cash = cash - $1 WHERE user_uuid = $2 RETURNING cash`

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var cash float32
	if err = tx.QueryRow(query, req.Cash, req.UserUUID).Scan(&cash); err != nil {
		tx.Rollback()
		return 0, err
	}

	return cash, nil
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
