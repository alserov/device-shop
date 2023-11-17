package postgres

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/order-service/internal/utils"
	"github.com/alserov/device-shop/order-service/pkg/entity"
	"github.com/alserov/device-shop/proto/gen"
	"sync"
	"time"
)

type Repo interface {
	CreateOrder(context.Context, *entity.CreateOrderReqWithDevices, *sql.Tx) error
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

func (r *repo) CreateOrder(ctx context.Context, req *entity.CreateOrderReqWithDevices, tx *sql.Tx) error {
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
				return
			}

			query = `UPDATE devices SET amount = amount - $1 WHERE uuid = $2`
			_, err = tx.Exec(query, device.Amount, device.UUID)
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
	query := `SELECT * FROM orders WHERE order_uuid = $1`

	rows, err := r.db.Query(query, orderUUID)
	if err != nil {
		return nil, err
	}

	var (
		devices    []*pb.Device
		createdAt  *time.Time
		statusCode = int32(-1)
	)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	chErr := make(chan error)

	for rows.Next() {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var orderedDevice entity.OrderedDevice
			if err = rows.Scan(&orderedDevice); err != nil {
				chErr <- err
			}
			if statusCode == -1 {
				statusCode = orderedDevice.Status
			}
			if createdAt == nil {
				createdAt = orderedDevice.CreatedAt
			}

			query = `SELECT * FROM devices WHERE device_uuid = $1`

			var device pb.Device

			if err = r.db.QueryRow(query, orderedDevice.DeviceUUID).Scan(&device); err != nil {
				chErr <- err
			}

			mu.Lock()
			devices = append(devices, &device)
			mu.Unlock()
		}()
	}

	go func() {
		wg.Wait()
		close(chErr)
	}()

	for e := range chErr {
		return &entity.CheckOrderRes{}, e
	}

	return &entity.CheckOrderRes{
		Devices:   devices,
		Status:    statusCode,
		CreatedAt: createdAt,
	}, err
}

func (r *repo) UpdateOrder(ctx context.Context, status string, orderUUID string) error {
	if utils.StatusToCode(status) == utils.CANCELED_CODE {
		query := `SELECT price FROM orders WHERE order_uuid = $1`
		rows, err := r.db.Query(query, orderUUID)
		if err != nil {
			return err
		}

		var (
			price    float32
			userUUID string
		)

		wg := &sync.WaitGroup{}
		mu := &sync.Mutex{}

		chErr := make(chan error, 1)
		for rows.Next() {
			wg.Add(1)
			go func() {
				defer wg.Done()

				var order entity.OrderedDevice
				if err = rows.Scan(&order); err != nil {
					chErr <- err
				}
				if userUUID == "" {
					userUUID = order.UserUUID
				}

				query = `SELECT price FROM devices WHERE device_uuid = $1`

				var devicePrice float32
				if err = r.db.QueryRow(query, order.DeviceUUID).Scan(&devicePrice); err != nil {
					chErr <- err
				}

				mu.Lock()
				price += devicePrice
				mu.Unlock()
			}()
		}
		wg.Wait()
		close(chErr)
		if err = <-chErr; err != nil {
			return err
		}

		query = `UPDATE users SET cash = cash + $1 WHERE user_uuid = $2`
		_, err = r.db.Exec(query, price, userUUID)
		if err != nil {
			return err
		}
	}

	query := `UPDATE orders SET status = $1 WHERE order_uuid = $2`

	_, err := r.db.Exec(query, utils.StatusToCode(status), orderUUID)
	if err != nil {
		return err
	}

	return nil
}
