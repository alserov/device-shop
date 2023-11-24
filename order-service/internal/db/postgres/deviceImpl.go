package postgres

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/order-service/internal/db"
	pb "github.com/alserov/device-shop/proto/gen"
	"sync"
)

func NewDeviceRepo(db *sql.DB) db.DeviceRepo {
	return &deviceRepo{
		db: db,
	}
}

type deviceRepo struct {
	db *sql.DB
}

func (r *deviceRepo) DecreaseDevicesAmount(_ context.Context, txCh chan<- *sql.Tx, devices []*pb.OrderDevice) error {
	query := `UPDATE devices SET amount = amount - $1 WHERE uuid = $2`
	var (
		chErr = make(chan error)
		wg    = &sync.WaitGroup{}
	)

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	txCh <- tx

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

func (r *deviceRepo) RollbackDevices(_ context.Context, txCh chan<- *sql.Tx, devices []*pb.Device) error {
	query := `UPDATE devices SET amount = amount + $1 WHERE uuid = $2`
	var (
		chErr = make(chan error)
		wg    = &sync.WaitGroup{}
	)

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	txCh <- tx

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
