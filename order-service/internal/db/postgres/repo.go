package postgres

import (
	"context"
	"database/sql"
	"github.com/alserov/shop/order-service/internal/utils"
	"sync"
	"time"
)

type Repo interface {
	CreateOrder(ctx context.Context, req *CreateOrderReq) error
	CheckOrder(ctx context.Context, orderUUID string) (*CheckOrderRes, error)
	UpdateOrder(ctx context.Context, status string, orderUUID string) error
}

type repo struct {
	db *sql.DB
}

func New(db *sql.DB) Repo {
	return &repo{
		db: db,
	}
}

func (r *repo) CreateOrder(ctx context.Context, req *CreateOrderReq) error {
	query := `INSERT INTO orders (user_uuid,order_uuid,device_uuid,amount,status,created_at) VALUES($1,$2,$3,$4,$5)`

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(req.Devices) + 1)

	chErr := make(chan error)

	go func() {
		defer wg.Done()
		query = `UPDATE users SET cash = cash - $1 WHERE user_uuid = $2`
		_, err = tx.Exec(query, utils.CountOrderPrice(req.Devices), req.UserUUID)
		if err != nil {
			chErr <- err
		}
	}()

	for _, v := range req.Devices {
		v := v
		go func() {
			defer wg.Done()

			_, err = tx.Exec(query, req.UserUUID, req.OrderUUID, v.Price, v.UUID, v.Amount, req.Status, req.CreatedAt)
			if err != nil {
				chErr <- err
				return
			}

			query = `UPDATE devices SET amount = amount - $1 WHERE uuid = $2`
			_, err = tx.Exec(query, v.Amount, v.UUID)
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

func (r *repo) CheckOrder(ctx context.Context, orderUUID string) (*CheckOrderRes, error) {
	query := `SELECT * FROM orders WHERE order_uuid = $1`

	rows, err := r.db.Query(query, orderUUID)
	if err != nil {
		return nil, err
	}

	var devices []*models.Device

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	chErr := make(chan error)

	for rows.Next() {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var device models.Device
			if err = rows.Scan(&device); err != nil {
				chErr <- err
			}

			mu.Lock()
			devices = append(devices, device)
			mu.Unlock()
		}()
	}

	go func() {
		wg.Wait()
		close(chErr)
	}()

	for e := range chErr {
		return &CheckOrderRes{}, e
	}

	return &CheckOrderRes{
		Devices:   devices,
		UserUUID:  devices[0].UserUUID,
		Status:    devices[0].Status,
		Price:     utils.CountOrderPrice(devices),
		CreatedAt: devices[0].CreatedAt,
	}, err
}

func (r *repo) UpdateOrder(ctx context.Context, status string, orderUUID string) error {
	if utils.StatusToCode(status) == utils.CANCELED_CODE {
		query := `SELECT price FROM orders WHERE order_uuid = $1`
		rows, err := r.db.Query(query, orderUUID)
		if err != nil {
			return err
		}

		var devices []*models.Device

		wg := &sync.WaitGroup{}
		mu := &sync.Mutex{}

		chErr := make(chan error, 1)
		for rows.Next() {
			wg.Add(1)
			go func() {
				defer wg.Done()

				var device models.Device
				if err = rows.Scan(&device); err != nil {
					chErr <- err
				}

				mu.Lock()
				devices = append(devices, device)
				mu.Unlock()
			}()
		}
		wg.Wait()
		close(chErr)
		if err = <-chErr; err != nil {
			return err
		}

		query = `UPDATE users SET cash = cash + $1 WHERE user_uuid = $2`
		_, err = r.db.Exec(query, utils.CountOrderPrice(devices), devices[0].UserUUID)
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

type CreateOrderReq struct {
	UserUUID  string
	Devices   []*models.Device
	OrderUUID string
	Status    int
	CreatedAt *time.Time
}

type CheckOrderRes struct {
	UserUUID  string           `bson:"userUUID"`
	Devices   []*models.Device `bson:"devices"`
	Status    int              `bson:"status"`
	Price     uint             `bson:"price"`
	CreatedAt *time.Time       `bson:"createdAt"`
}
