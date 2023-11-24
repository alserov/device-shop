package db

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/order-service/internal/entity"
	pb "github.com/alserov/device-shop/proto/gen"
)

type OrderRepo interface {
	CreateOrder(ctx context.Context, txCh chan<- *sql.Tx, req *pb.CreateOrderReq, info *entity.OrderAdditional) error
	CheckOrder(ctx context.Context, orderUUID string) (*entity.CheckOrderRes, error)
	UpdateOrder(ctx context.Context, status string, orderUUID string) error
}
