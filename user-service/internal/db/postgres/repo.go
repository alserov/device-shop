package postgres

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/user-service/internal/entity"
)

type Repository interface {
	GetInfo(context.Context, string) (*entity.RepoGetInfoRes, error)
	TopUpBalance(context.Context, *entity.TopUpBalanceReq) (float32, error)
	DebitBalance(context.Context, *entity.DebitBalanceReq) (float32, error)
}

func NewRepo(db *sql.DB) Repository {
	return &repo{
		db: db,
	}
}

type repo struct {
	db *sql.DB
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
	query := `UPDATE users SET cash = cash - $1 WHERE uuid = $2 RETURNING cash`

	var cash float32
	if err := r.db.QueryRow(query, req.Cash, req.UserUUID).Scan(&cash); err != nil {
		return 0, err
	}

	return cash, nil
}
