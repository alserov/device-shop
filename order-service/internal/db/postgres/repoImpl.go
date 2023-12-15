package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	oStatus "github.com/alserov/device-shop/order-service/internal/utils/status"

	"github.com/alserov/device-shop/order-service/internal/db"
	"github.com/alserov/device-shop/order-service/internal/db/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"sync"
)

func NewRepo(db *sql.DB, log *slog.Logger) db.OrderRepo {
	return &repo{
		log: log,
		db:  db,
	}
}

type repo struct {
	log *slog.Logger
	db  *sql.DB
}

const (
	internalError = "internal error"
	orderNotFound = "order not found"
)

func (r *repo) CreateOrderTx(_ context.Context, req models.CreateOrderReq) (*sql.Tx, error) {
	op := "repo.CreateOrder"
	var (
		chErr = make(chan error)
		wg    = &sync.WaitGroup{}
	)
	wg.Add(len(req.OrderDevices) + 1)

	tx, err := r.db.Begin()
	if err != nil {
		r.log.Error("failed to start transaction", slog.String("error", err.Error()), slog.String("op", op))
		return tx, err
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
		return tx, status.Error(codes.Internal, internalError)
	}

	return tx, nil
}

func (r *repo) GetOrderDevices(_ context.Context, orderUUID string) ([]models.OrderDevice, error) {
	query := `SELECT device_uuid, amount FROM ordered_devices WHERE order_uuid = $1`

	var orderedDevices []models.OrderDevice

	rows, err := r.db.Query(query, orderUUID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("there are no devices ordered with uuid: %s", orderUUID))
	}
	if err != nil {
		return nil, status.Error(codes.Internal, internalError)
	}

	for rows.Next() {
		var d models.OrderDevice
		if err = rows.Scan(&d); err != nil {
			return nil, status.Error(codes.Internal, internalError)
		}
		orderedDevices = append(orderedDevices, d)
	}

	return orderedDevices, nil
}

func (r *repo) CheckOrder(_ context.Context, orderUUID string) (models.CheckOrderRes, error) {
	op := "repo.CheckOrder"

	var (
		chErr   = make(chan error)
		wg      = &sync.WaitGroup{}
		devices []models.OrderDevice
		order   = models.Order{}
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
			devices = append(devices, models.OrderDevice{
				DeviceUUID: uuid,
				Amount:     amount,
			})
		}
	}()

	go func() {
		defer wg.Done()
		query := `SELECT total_price,status,created_at, user_uuid FROM orders WHERE order_uuid = $1`

		if err := r.db.QueryRow(query, orderUUID).Scan(&order.Price, &order.Status, &order.CreatedAt, &order.UserUUID); err != nil {
			chErr <- err
		}
	}()

	go func() {
		wg.Wait()
		close(chErr)
	}()

	for err := range chErr {
		if errors.Is(sql.ErrNoRows, err) {
			return models.CheckOrderRes{}, status.Error(codes.InvalidArgument, orderNotFound)
		}
		r.log.Error("failed to create order", slog.String("error", err.Error()), slog.String("op", op))
		return models.CheckOrderRes{}, err
	}

	return models.CheckOrderRes{
		OrderDevices: devices,
		Status:       order.Status,
		CreatedAt:    order.CreatedAt,
		OrderPrice:   order.Price,
		UserUUID:     order.UserUUID,
	}, nil
}

func (r *repo) CancelOrderTx(ctx context.Context, orderUUID string) (models.CancelOrderRes, *sql.Tx, error) {
	op := "repo.CancelOrderTx"
	query := `UPDATE orders SET status = $1 WHERE order_uuid = $2 RETURNING order_price,user_uuid`

	tx, err := r.db.Begin()
	if err != nil {
		return models.CancelOrderRes{}, tx, status.Error(codes.Internal, internalError)
	}

	var res models.CancelOrderRes
	err = tx.QueryRow(query, oStatus.CANCELED_CODE, orderUUID).Scan(&res)
	if errors.Is(err, sql.ErrNoRows) {
		return models.CancelOrderRes{}, tx, status.Error(codes.NotFound, orderNotFound)
	}
	if err != nil {
		r.log.Error("failed to execute query", slog.String("error", err.Error()), slog.String("op", op))
		return models.CancelOrderRes{}, tx, status.Error(codes.Internal, internalError)
	}

	return res, tx, nil
}

func (r *repo) UpdateOrder(_ context.Context, orderStatus string, orderUUID string) error {
	op := "repo.UpdateOrder"
	var (
		wg    = &sync.WaitGroup{}
		chErr = make(chan error)
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		query := `UPDATE orders SET status = $1 WHERE order_uuid = $2`

		_, err := r.db.Exec(query, oStatus.StatusToCode(orderStatus), orderUUID)
		if err != nil {
			r.log.Error("failed to delete order from orders", slog.String("error", err.Error()), slog.String("op", op))
			chErr <- err
		}
	}()

	go func() {
		defer wg.Done()
		query := `DELETE FROM order_devices * WHERE order_uuid =$1`

		_, err := r.db.Exec(query, orderUUID)
		if err != nil {
			r.log.Error("failed to delete order devices from order_devices", slog.String("error", err.Error()), slog.String("op", op))
			chErr <- err
		}
	}()

	go func() {
		wg.Wait()
		close(chErr)
	}()

	for _ = range chErr {
		return status.Error(codes.Internal, internalError)
	}

	return nil
}
