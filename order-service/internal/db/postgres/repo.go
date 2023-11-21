package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alserov/device-shop/order-service/internal/entity"
	"github.com/alserov/device-shop/order-service/internal/utils"
	pb "github.com/alserov/device-shop/proto/gen"
	"sync"
)

type Repository interface {
	CreateOrder(ctx context.Context, tx *sql.Tx, req *pb.CreateOrderReq, info *entity.OrderAdditional) error
	CheckOrder(ctx context.Context, orderUUID string) (*entity.CheckOrderRes, error)
	UpdateOrder(ctx context.Context, status string, orderUUID string) error

	DecreaseDevicesAmount(ctx context.Context, tx *sql.Tx, devices []*pb.OrderDevice) error
	DebitBalance(ctx context.Context, tx *sql.Tx, userUUID string, cash float32) error

	RollbackDevices(ctx context.Context, tx *sql.Tx, devices []*pb.Device) error
	RollbackBalance(ctx context.Context, tx *sql.Tx, userUUID string, orderUUID string, cash float32) error

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
	var (
		chErr = make(chan error)
		wg    = &sync.WaitGroup{}
	)
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

		rows, err := r.orders.Query(query, orderUUID)
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

		if err := r.orders.QueryRow(query, orderUUID).Scan(&order.TotalPrice, &order.Status, &order.CreatedAt, &order.UserUUID); err != nil {
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

func (r *repo) UpdateOrder(ctx context.Context, status string, orderUUID string) error {
	query := `UPDATE orders SET status = $1 WHERE order_uuid = $2`

	_, err := r.orders.Exec(query, utils.StatusToCode(status), orderUUID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) DecreaseDevicesAmount(ctx context.Context, tx *sql.Tx, devices []*pb.OrderDevice) error {
	query := `UPDATE devices SET amount = amount - $1 WHERE uuid = $2`
	var (
		chErr = make(chan error)
		wg    = &sync.WaitGroup{}
	)

	for _, d := range devices {
		d := d
		go func() {
			if _, err := tx.Exec(query, d.Amount, d.DeviceUUID); err != nil {
				chErr <- err
			}
		}()
	}

	go func() {
		wg.Wait()
		close(chErr)
	}()

	for err := range chErr {
		return err
	}

	return nil
}

func (r *repo) DebitBalance(ctx context.Context, tx *sql.Tx, userUUID string, cash float32) error {
	query := `UPDATE users SET cash = cash - $1 WHERE uuid = $2`

	if _, err := tx.Exec(query, cash, userUUID); err != nil {
		return err
	}
	return nil
}

func (r *repo) RollbackDevices(ctx context.Context, tx *sql.Tx, devices []*pb.Device) error {
	query := `UPDATE devices SET amount = amount + $1 WHERE uuid = $2`
	var (
		chErr = make(chan error)
		wg    = &sync.WaitGroup{}
	)

	for _, d := range devices {
		d := d
		go func() {
			if _, err := tx.Exec(query, d.Amount, d.UUID); err != nil {
				chErr <- err
			}
		}()
	}

	go func() {
		wg.Wait()
		close(chErr)
	}()

	for err := range chErr {
		return err
	}

	return nil
}

func (r *repo) RollbackBalance(ctx context.Context, tx *sql.Tx, userUUID string, orderUUID string, cash float32) error {
	query := `SELECT status FROM orders WHERE order_uuid = $1`

	var status int32
	if err := r.orders.QueryRow(query, orderUUID).Scan(&status); err != nil {
		return err
	}
	if status == utils.CANCELED_CODE {
		return errors.New("this order is already canceled")
	}

	query = `UPDATE users SET cash = cash + $1 WHERE uuid = $2`

	if _, err := tx.Exec(query, cash, userUUID); err != nil {
		return err
	}
	return nil
}
