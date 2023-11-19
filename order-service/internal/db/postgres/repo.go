package postgres

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/order-service/internal/utils"
	"github.com/alserov/device-shop/order-service/pkg/entity"
	"sync"
	"time"
)

type Repo interface {
	CreateOrder(context.Context, *entity.CreateOrderReqWithDevices) error
	CheckOrder(context.Context, string) (*entity.CheckOrderRes, error)
	UpdateOrder(context.Context, string, string) error
	GetDB() *sql.DB
}

func New(db *sql.DB) Repo {
	return &repo{
		db: db,
	}
}

type repo struct {
	db *sql.DB
}

func (r *repo) GetDB() *sql.DB {
	return r.db
}

func (r *repo) CreateOrder(ctx context.Context, req *entity.CreateOrderReqWithDevices) error {
	query := `INSERT INTO orders (user_uuid,order_uuid,device_uuid,amount,status,created_at) VALUES($1,$2,$3,$4,$5,$6)`

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(req.Devices))

	chErr := make(chan error)

	for _, device := range req.Devices {
		device := device
		go func() {
			defer wg.Done()
			_, err = tx.Exec(query, req.UserUUID, req.OrderUUID, device.UUID, device.Amount, req.Status, req.CreatedAt)
			if err != nil {
				chErr <- err
			}
		}()
	}

	go func() {
		wg.Wait()
		close(chErr)
		tx.Commit()
	}()

	for e := range chErr {
		tx.Rollback()
		return e
	}

	return nil
}

func (r *repo) CheckOrder(ctx context.Context, orderUUID string) (*entity.CheckOrderRes, error) {
	query := `SELECT device_uuid, amount, status, created_at FROM orders WHERE order_uuid = $1`

	rows, err := r.db.Query(query, orderUUID)
	if err != nil {
		return nil, err
	}

	var (
		devices    []*entity.OrderDevice
		createdAt  *time.Time
		statusCode = int32(-1)
	)

	for rows.Next() {
		var orderedDevice entity.OrderedDevice
		if err = rows.Scan(&orderedDevice.DeviceUUID, &orderedDevice.Amount, &orderedDevice.Status, &orderedDevice.CreatedAt); err != nil {
			return &entity.CheckOrderRes{}, err
		}
		if statusCode == -1 {
			statusCode = orderedDevice.Status
		}
		if createdAt == nil {
			createdAt = orderedDevice.CreatedAt
		}
		devices = append(devices, &entity.OrderDevice{
			DeviceUUID: orderedDevice.DeviceUUID,
			Amount:     orderedDevice.Amount,
		})
	}

	return &entity.CheckOrderRes{
		Devices:   devices,
		Status:    statusCode,
		CreatedAt: createdAt,
	}, err
}

func (r *repo) UpdateOrder(ctx context.Context, status string, orderUUID string) error {
	query := `UPDATE orders SET status = $1 WHERE order_uuid = $2`

	_, err := r.db.Exec(query, utils.StatusToCode(status), orderUUID)
	if err != nil {
		return err
	}

	return nil
}
