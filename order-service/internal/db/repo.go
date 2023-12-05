package db

import (
	"context"
	"github.com/alserov/device-shop/order-service/internal/db/models"
)

type OrderRepo interface {
	CreateOrder(ctx context.Context, req models.CreateOrderReq) error
	CheckOrder(ctx context.Context, orderUUID string) (models.CheckOrderRes, error)
	UpdateOrder(ctx context.Context, status string, orderUUID string) error
}
