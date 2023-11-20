package postgres

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/order-service/internal/entity"
	"github.com/alserov/device-shop/order-service/internal/utils"
	pb "github.com/alserov/device-shop/proto/gen"
	"sync"
	"time"
)

type Repository interface {
	CreateOrder(ctx context.Context, tx *sql.Tx, req *pb.CreateOrderReq, info *entity.OrderAdditional) error
	CheckOrder(ctx context.Context, orderUUID string) (*entity.CheckOrderRes, error)
	UpdateOrder(ctx context.Context, status string, orderUUID string) error

	DecreaseDevicesAmount(ctx context.Context, tx *sql.Tx, devices []*pb.OrderDevice) error
	DebitBalance(ctx context.Context, tx *sql.Tx, userUUID string, cash float32) error

	RollbackDevices(ctx context.Context)
	RollbackBalance(ctx context.Context)

	GetOrdersDB() *sql.DB
	GetUsersDB() *sql.DB
	GetDevicesDB() *sql.DB
}

func New(ordersDB, devicesDB, usersDB *sql.DB) Repository {
	return &repo{
		orders:  ordersDB,
		devices: devicesDB,
		users:   usersDB,
	}
}

type repo struct {
	orders  *sql.DB
	devices *sql.DB
	users   *sql.DB
}

func (r *repo) GetOrdersDB() *sql.DB {
	return r.orders
}

func (r *repo) GetUsersDB() *sql.DB {
	return r.users
}

func (r *repo) GetDevicesDB() *sql.DB {
	return r.devices
}

func (r *repo) CreateOrder(ctx context.Context, tx *sql.Tx, req *pb.CreateOrderReq, info *entity.OrderAdditional) error {
	chErr := make(chan error)
	wg := &sync.WaitGroup{}
	wg.Add(len(req.Devices) + 1)

	go func() {
		defer wg.Done()
		query := `INSERT INTO orders (order_uuid,user_uuid,total_price,status,created_at) VALUES($1,$2,$3,$4,$5)`
		_, err := tx.Exec(query, info.OrderUUID, req.UserUUID, info.TotalPrice, info.Status, info.CreatedAt)
		if err != nil {
			chErr <- err
		}
	}()

	go func() {
		query := `INSERT INTO ordered_devices (order_uuid, device_uuid, amount) VALUES($1,$2,$3)`
		for _, device := range req.Devices {
			device := device
			go func() {
				defer wg.Done()
				_, err := tx.Exec(query, info.OrderUUID, device.DeviceUUID, device.Amount)
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

	for err := range chErr {
		return err
	}

	return nil
}

func (r *repo) CheckOrder(ctx context.Context, orderUUID string) (*entity.CheckOrderRes, error) {
	query := `SELECT device_uuid, amount, status, created_at FROM orders WHERE order_uuid = $1`

	rows, err := r.orders.Query(query, orderUUID)
	if err != nil {
		return nil, err
	}

	var (
		devices    []*pb.Device
		createdAt  *time.Time
		statusCode = int32(-1)
	)

	for rows.Next() {
		var orderedDevice entity.OrderDevice
		if err = rows.Scan(&orderedDevice.UUID, &orderedDevice.Amount, &orderedDevice.Status, &orderedDevice.CreatedAt); err != nil {
			return &entity.CheckOrderRes{}, err
		}
		if statusCode == -1 {
			statusCode = orderedDevice.Status
		}
		if createdAt == nil {
			createdAt = orderedDevice.CreatedAt
		}
		devices = append(devices, &pb.Device{
			UUID:   orderedDevice.UUID,
			Amount: orderedDevice.Amount,
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

	_, err := r.orders.Exec(query, utils.StatusToCode(status), orderUUID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) TakeDevices(ctx context.Context, tx *sql.Tx, devices []*pb.OrderDevice) ([]*pb.Device, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repo) DebitBalance(ctx context.Context, tx *sql.Tx, userUUID string, cash float32) error {
	//TODO implement me
	panic("implement me")
}
