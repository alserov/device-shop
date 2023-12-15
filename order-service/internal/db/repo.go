package db

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/order-service/internal/db/models"
)

type OrderRepo interface {
	CreateOrderTx(ctx context.Context, req models.CreateOrderReq) (*sql.Tx, error)
	CheckOrder(ctx context.Context, orderUUID string) (models.CheckOrderRes, error)
	UpdateOrder(ctx context.Context, status string, orderUUID string) error
	CancelOrderTx(ctx context.Context, orderUUID string) (models.CancelOrderRes, *sql.Tx, error)
	CancelOrderDevicesTx(ctx context.Context, orderUUID string) ([]models.OrderDevice, *sql.Tx, error)
}
