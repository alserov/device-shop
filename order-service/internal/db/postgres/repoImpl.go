package postgres

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/order-service/internal/db"
	"github.com/alserov/device-shop/order-service/internal/db/models"
	"github.com/alserov/device-shop/order-service/internal/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"sync"
)

func NewOrderRepo(db *sql.DB) db.OrderRepo {
	return &repo{
		db: db,
	}
}

type repo struct {
	log *slog.Logger
	db  *sql.DB
}

const (
	internalError = "internal error"
)

func (r *repo) CreateOrder(_ context.Context, req models.CreateOrderReq) error {
	op := "repo.CreateOrder"
	var (
		chErr = make(chan error)
		wg    = &sync.WaitGroup{}
	)
	wg.Add(len(req.OrderDevices) + 1)

	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		r.log.Error("failed to start transaction", slog.String("error", err.Error()), slog.String("op", op))
		return err
	}

	go func() {
		defer wg.Done()
		query := `INSERT INTO orders (order_uuid,user_uuid,total_price,status,created_at) VALUES($1,$2,$3,$4,$5)`
		_, err := tx.Exec(query, req.OrderUUID, req.UserUUID, req.OrderPrice, req.Status, req.CreatedAt)
		if err != nil {
			chErr <- err
		}
	}()

	go func() {
		query := `INSERT INTO ordered_devices (order_uuid, device_uuid, amount) VALUES($1,$2,$3)`
		for _, device := range req.OrderDevices {
			device := device
			go func() {
				defer wg.Done()
				_, err := tx.Exec(query, req.OrderUUID, device.DeviceUUID, device.Amount)
				if err != nil {
					chErr <- err
				}
			}()
		}
	}()

	go func() {
		wg.Wait()
		close(chErr)
	}()

	for err = range chErr {
		tx.Rollback()
		r.log.Error("failed to create order", slog.String("error", err.Error()), slog.String("op", op))
		return status.Error(codes.Internal, internalError)
	}

	tx.Commit()

	return nil
}

func (r *orderRepo) CheckOrder(_ context.Context, orderUUID string) (*entity.CheckOrderRes, error) {
	var (
		chDevices = make(chan *pb.Device)
		chErr     = make(chan error)
		wg        = &sync.WaitGroup{}
		devices   []*pb.Device
		order     = &entity.CheckOrderRes{}
	)
	wg.Add(2)

	go func() {
		defer wg.Done()
		query := `SELECT device_uuid,amount FROM ordered_devices WHERE order_uuid = $1`

		rows, err := r.db.Query(query, orderUUID)
		if err != nil {
			chErr <- err
		}

		for rows.Next() {
			var (
				uuid   string
				amount uint32
			)
			if err = rows.Scan(&uuid, &amount); err != nil {
				chErr <- err
			}
			chDevices <- &pb.Device{
				UUID:   uuid,
				Amount: amount,
			}
		}
	}()

	go func() {
		defer wg.Done()
		query := `SELECT total_price,status,created_at, user_uuid FROM orders WHERE order_uuid = $1`

		if err := r.db.QueryRow(query, orderUUID).Scan(&order.TotalPrice, &order.Status, &order.CreatedAt, &order.UserUUID); err != nil {
			chErr <- err
		}
	}()

	go func() {
		for d := range chDevices {
			devices = append(devices, d)
		}
	}()

	go func() {
		wg.Wait()
		close(chDevices)
		close(chErr)
	}()

	for err := range chErr {
		return nil, err
	}

	return &entity.CheckOrderRes{
		Devices:    devices,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt,
		TotalPrice: order.TotalPrice,
		UserUUID:   order.UserUUID,
		UUID:       orderUUID,
	}, nil
}

func (r *orderRepo) UpdateOrder(_ context.Context, status string, orderUUID string) error {
	query := `UPDATE orders SET status = $1 WHERE order_uuid = $2`

	_, err := r.db.Exec(query, status.StatusToCode(status), orderUUID)
	if err != nil {
		return err
	}

	return nil
}
