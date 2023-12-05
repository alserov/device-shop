package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alserov/device-shop/proto/gen/order"

	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/order-service/internal/db"
	"github.com/alserov/device-shop/order-service/internal/db/postgres"
	"github.com/alserov/device-shop/order-service/internal/service/models"
	"github.com/alserov/device-shop/order-service/internal/utils/converter"
	"github.com/alserov/device-shop/order-service/internal/utils/status"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sync"
)

type Service interface {
	CreateOrder(ctx context.Context, req models.CreateOrderReq) (models.CreateOrderRes, error)
	CheckOrder(ctx context.Context, req models.CheckOrderReq) (models.CheckOrderRes, error)
}

type service struct {
	db db.OrderRepo

	conv *converter.ServiceConverter
}

func NewService(ordersDB *sql.DB) Service {
	return &service{
		db: postgres.NewOrderRepo(ordersDB),
	}
}

func (s *service) CreateOrder(ctx context.Context, req models.CreateOrderReq) (models.CreateOrderRes, error) {
	orderUUID := uuid.New().String()

	if err := s.db.CreateOrder(ctx, s.conv.CreateOrderReqToRepo(req, orderUUID)); err != nil {
		return models.CreateOrderRes{}, err
	}

	return s.conv.CreateOrderResToService(orderUUID), nil
}

func (s *service) CheckOrder(ctx context.Context, req models.CheckOrderReq) (models.CheckOrderRes, error) {
	order, err := s.db.CheckOrder(ctx, req.OrderUUID)
	if err != nil {
		return models.CheckOrderRes{}, err
	}
	return s.conv.CheckOrderToService(order), nil
}

func (s *service) UpdateOrder(ctx context.Context, req *order.UpdateOrderReq) (*order.UpdateOrderRes, error) {
	if status.StatusToCode(req.Status) == status.CANCELED_CODE {
		order, err := s.db.CheckOrder(ctx, req.OrderUUID)
		if err != nil {
			return nil, err
		}

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		var (
			chErr = make(chan error)
			wg    = &sync.WaitGroup{}
		)
		wg.Add(txTwo)
		txCh := make(chan *sql.Tx, txTwo)

		go func() {
			defer wg.Done()
			if err = s.device.RollbackDevices(ctx, txCh, order.Devices); err != nil {
				chErr <- err
			}
		}()

		go func() {
			defer wg.Done()
			o, err := s.order.CheckOrder(ctx, req.OrderUUID)
			if o.Status == status.CANCELED_CODE {
				chErr <- errors.New("order is already canceled")
				return
			}
			if err = s.user.RollbackBalance(ctx, txCh, order.UserUUID, order.UUID, order.TotalPrice); err != nil {
				chErr <- err
			}
		}()

		go func() {
			wg.Wait()
			close(chErr)
			close(txCh)
		}()

		for err = range chErr {
			for tx := range txCh {
				tx.Rollback()
			}
			return &pb.UpdateOrderRes{}, err
		}

		for tx := range txCh {
			tx.Commit()
		}
	}
	if err := s.db.UpdateOrder(ctx, req.Status, req.OrderUUID); err != nil {
		return nil, err
	}
	return , nil
}
