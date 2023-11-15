package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alserov/device-shop/device-service/pkg/entity"
)

type Repository interface {
	CreateDevice(ctx context.Context, device *entity.Device) error
	DeleteDevice(ctx context.Context, uuid string) error
	UpdateDevice(ctx context.Context, device *entity.UpdateDeviceReq) error
	GetAllDevices(ctx context.Context, index uint32, amount uint32) ([]*entity.Device, error)
	GetDevicesByTitle(ctx context.Context, title string) ([]*entity.Device, error)
	GetDeviceByUUID(ctx context.Context, uuid string) (*entity.Device, error)
	GetDevicesByManufacturer(ctx context.Context, manu string) ([]*entity.Device, error)
	GetDevicesByPrice(ctx context.Context, min uint, max uint) ([]*entity.Device, error)
}

type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) CreateDevice(ctx context.Context, device *entity.Device) error {
	query := `INSERT INTO devices (uuid, title, description, price, manufacturer) VALUES($1,$2,$3,$4,$5)`

	_, err := r.db.Exec(query, device.UUID, device.Title, device.Description, device.Price, device.Manufacturer)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) DeleteDevice(ctx context.Context, uuid string) error {
	query := `DELETE FROM devices WHERE uuid = $1`

	_, err := r.db.Exec(query, uuid)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) UpdateDevice(ctx context.Context, device *entity.UpdateDeviceReq) error {
	query := `UPDATE devices SET title = $1, description = $2, price = $3 WHERE uuid = $4`

	_, err := r.db.Exec(query, device.Title, device.Description, device.Price, device.UUID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetAllDevices(ctx context.Context, index uint32, amount uint32) ([]*entity.Device, error) {
	query := `SELECT * FROM devices LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, amount, index)
	if err != nil {
		return nil, nil
	}

	devices := make([]*entity.Device, 0, amount)
	for rows.Next() {
		d := entity.Device{}
		if err = rows.Scan(&d.UUID, &d.Title, &d.Description, &d.Price, &d.Manufacturer); err != nil {
			return nil, err
		}
		devices = append(devices, &d)
	}

	return devices, nil
}

func (r *repo) GetDevicesByTitle(ctx context.Context, title string) ([]*entity.Device, error) {
	query := `SELECT * FROM devices WHERE title LIKE $1`

	rows, err := r.db.Query(query, "%"+title+"%")
	if err != nil {
		return nil, err
	}

	var devices []*entity.Device
	for rows.Next() {
		d := entity.Device{}
		if err = rows.Scan(&d.UUID, &d.Title, &d.Description, &d.Price, &d.Manufacturer); err != nil {
			return nil, err
		}
		devices = append(devices, &d)
	}

	return devices, nil
}

func (r *repo) GetDeviceByUUID(ctx context.Context, uuid string) (*entity.Device, error) {
	query := `SELECT * FROM devices WHERE uuid = $1`

	d := entity.Device{}

	err := r.db.QueryRow(query, uuid).Scan(&d.UUID, &d.Title, &d.Description, &d.Price, &d.Manufacturer)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return &d, nil
}

func (r *repo) GetDevicesByManufacturer(ctx context.Context, manu string) ([]*entity.Device, error) {
	query := `SELECT * FROM devices WHERE manufacturer = $1`

	rows, err := r.db.Query(query, manu)
	if err != nil {
		return nil, err
	}

	var devices []*entity.Device
	for rows.Next() {
		d := entity.Device{}
		if err = rows.Scan(&d.UUID, &d.Title, &d.Description, &d.Price, &d.Manufacturer); err != nil {
			return nil, err
		}
		devices = append(devices, &d)
	}

	return devices, nil
}

func (r *repo) GetDevicesByPrice(ctx context.Context, min uint, max uint) ([]*entity.Device, error) {
	query := `SELECT * FROM devices WHERE price BETWEEN $1 AND $2`

	rows, err := r.db.Query(query, min, max)
	if err != nil {
		return nil, err
	}

	var devices []*entity.Device
	for rows.Next() {
		d := entity.Device{}
		if err = rows.Scan(&d.UUID, &d.Title, &d.Description, &d.Price, &d.Manufacturer); err != nil {
			return nil, err
		}
		devices = append(devices, &d)
	}

	return devices, nil
}
