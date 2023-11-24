package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/alserov/device-shop/order-service/internal/db"
)

func NewUserRepo(db *sql.DB) db.UserRepo {
	return &userRepo{
		db: db,
	}
}

type userRepo struct {
	db *sql.DB
}

func (r *userRepo) DebitBalance(_ context.Context, txCh chan<- *sql.Tx, userUUID string, cash float32) error {
	query := `UPDATE users SET cash = cash - $1 WHERE uuid = $2`

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	txCh <- tx

	if _, err := tx.Exec(query, cash, userUUID); err != nil {
		return fmt.Errorf("not enough balance")
	}
	return nil
}

func (r *userRepo) RollbackBalance(_ context.Context, txCh chan<- *sql.Tx, userUUID string, orderUUID string, cash float32) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	txCh <- tx

	query := `UPDATE users SET cash = cash + $1 WHERE uuid = $2`

	if _, err := tx.Exec(query, cash, userUUID); err != nil {
		return err
	}
	return nil
}
