package db

import (
	"context"
	"database/sql"
)

type UserRepo interface {
	DebitBalance(ctx context.Context, txCh chan<- *sql.Tx, userUUID string, cash float32) error
	RollbackBalance(ctx context.Context, txCh chan<- *sql.Tx, userUUID string, orderUUID string, cash float32) error
}
