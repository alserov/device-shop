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

func NewRepo(db *sql.DB, log *slog.Logger) db.Repository {
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
		if err = tx.Rollback(); err != nil {
			r.log.Error("failed to rollback", slog.String("error", err.Error()), slog.String("op", op))
			return tx, status.Error(codes.Internal, internalError)
		}
		r.log.Error("failed to create order", slog.String("error", err.Error()), slog.String("op", op))
		return tx, err
	}

	return tx, nil
}

func (r *repo) CancelOrderDevicesTx(_ context.Context, orderUUID string) (db.Tx, error) {
	query := `DELETE FROM ordered_devices WHERE order_uuid = $1 RETURNING device_uuid, amount`

	tx, err := r.db.Begin()
	if err != nil {
		return nil, status.Error(codes.Internal, internalError)
	}

	var devices []models.OrderDevice

	rows, err := tx.Query(query, orderUUID)
	if errors.Is(err, sql.ErrNoRows) {
		return &db.CancelOrderDevices{
			Devices: devices,
			Tx:      tx,
		}, status.Error(codes.NotFound, fmt.Sprintf("there are no devices ordered with uuid: %s", orderUUID))
	}
	if err != nil {
		return nil, status.Error(codes.Internal, internalError)
	}

	for rows.Next() {
		var d models.OrderDevice
		if err = rows.Scan(&d); err != nil {
			return nil, status.Error(codes.Internal, internalError)
		}
		devices = append(devices, d)
	}

	return &db.CancelOrderDevices{
		Devices: devices,
		Tx:      tx,
	}, nil
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
			err = rows.Scan(&uuid, &amount)
			if errors.Is(sql.ErrNoRows, err) {
				r.log.Error("failed to get order device", slog.String("error", err.Error()), slog.String("op", op))
				chErr <- status.Error(codes.InvalidArgument, orderNotFound)
			}
			if err != nil {
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

		err := r.db.QueryRow(query, orderUUID).Scan(&order.Price, &order.Status, &order.CreatedAt, &order.UserUUID)
		if errors.Is(sql.ErrNoRows, err) {
			r.log.Error("failed to get order", slog.String("error", err.Error()), slog.String("op", op))
			chErr <- status.Error(codes.InvalidArgument, orderNotFound)
		}
		if err != nil {
			chErr <- err
		}
	}()

	go func() {
		wg.Wait()
		close(chErr)
	}()

	for err := range chErr {
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

func (r *repo) CancelOrderTx(_ context.Context, orderUUID string) (db.Tx, error) {
	op := "repo.CancelOrderTx"
	query := `UPDATE orders SET status = $1 WHERE order_uuid = $2 RETURNING total_price,user_uuid`

	var res db.CancelOrder

	tx, err := r.db.Begin()
	res.Tx = tx
	if err != nil {
		return &res, status.Error(codes.Internal, internalError)
	}

	err = tx.QueryRow(query, oStatus.CANCELED_CODE, orderUUID).Scan(&res.Price, &res.UserUUID)
	if errors.Is(err, sql.ErrNoRows) {
		return &db.CancelOrder{}, status.Error(codes.NotFound, orderNotFound)
	}
	if err != nil {
		r.log.Error("failed to execute query", slog.String("error", err.Error()), slog.String("op", op))
		return &res, status.Error(codes.Internal, internalError)
	}

	return &res, nil
}

func (r *repo) UpdateOrder(_ context.Context, orderStatus string, orderUUID string) error {
	op := "repo.UpdateOrder"

	query := `UPDATE orders SET status = $1 WHERE order_uuid = $2`

	_, err := r.db.Exec(query, oStatus.StatusToCode(orderStatus), orderUUID)
	if errors.Is(err, sql.ErrNoRows) {
		return status.Error(codes.NotFound, orderNotFound)
	}
	if err != nil {
		r.log.Error("failed to delete order from orders", slog.String("error", err.Error()), slog.String("op", op))
		return status.Error(codes.Internal, internalError)
	}

	return nil
}
