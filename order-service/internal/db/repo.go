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
	CancelOrderTx(ctx context.Context, orderUUID string) (Tx, error)
	CancelOrderDevicesTx(ctx context.Context, orderUUID string) (Tx, error)
}

type Tx interface {
	Value() interface{}
	GetTx() SqlTx
}

type SqlTx interface {
	Commit() error
	Rollback() error
}

type CancelOrder struct {
	Price    float32
	UserUUID string
	Tx       SqlTx
}

func (cor *CancelOrder) Value() interface{} {
	return cor
}

func (cor *CancelOrder) GetTx() SqlTx {
	return cor.Tx
}

type CancelOrderDevices struct {
	Devices []models.OrderDevice
	Tx      SqlTx
}

func (codr *CancelOrderDevices) Value() interface{} {
	return codr
}

func (codr *CancelOrderDevices) GetTx() SqlTx {
	return codr.Tx
}
