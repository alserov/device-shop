package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alserov/device-shop/device-service/internal/db"
	pb "github.com/alserov/device-shop/proto/gen"
)

func NewDeviceRepo(db *sql.DB) db.DeviceRepo {
	return &deviceRepo{
		db: db,
	}
}

type deviceRepo struct {
	db *sql.DB
}

func (r *deviceRepo) GetAllDevices(_ context.Context, index uint32, amount uint32) ([]*pb.Device, error) {
	query := `SELECT * FROM devices LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, amount, index)
	if err != nil {
		return nil, nil
	}

	devices := make([]*pb.Device, 0, amount)
	for rows.Next() {
		d := pb.Device{}
		if err = rows.Scan(&d.UUID, &d.Title, &d.Description, &d.Price, &d.Manufacturer, &d.Amount); err != nil {
			return nil, err
		}
		devices = append(devices, &d)
	}

	return devices, nil
}

func (r *deviceRepo) GetDevicesByTitle(_ context.Context, title string) ([]*pb.Device, error) {
	query := `SELECT * FROM devices WHERE title LIKE $1`

	rows, err := r.db.Query(query, "%"+title+"%")
	if err != nil {
		return nil, err
	}

	var devices []*pb.Device
	for rows.Next() {
		d := pb.Device{}
		if err = rows.Scan(&d.UUID, &d.Title, &d.Description, &d.Price, &d.Manufacturer, &d.Amount); err != nil {
			return nil, err
		}
		devices = append(devices, &d)
	}

	return devices, nil
}

func (r *deviceRepo) GetDeviceByUUID(_ context.Context, uuid string) (*pb.Device, error) {
	query := `SELECT * FROM devices WHERE uuid = $1`

	d := pb.Device{}

	err := r.db.QueryRow(query, uuid).Scan(&d.UUID, &d.Title, &d.Description, &d.Price, &d.Manufacturer, &d.Amount)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return &d, nil
}

func (r *deviceRepo) GetDevicesByManufacturer(_ context.Context, manu string) ([]*pb.Device, error) {
	query := `SELECT * FROM devices WHERE manufacturer = $1`

	rows, err := r.db.Query(query, manu)
	if err != nil {
		return nil, err
	}

	var devices []*pb.Device
	for rows.Next() {
		d := pb.Device{}
		if err = rows.Scan(&d.UUID, &d.Title, &d.Description, &d.Price, &d.Manufacturer, &d.Amount); err != nil {
			return nil, err
		}
		devices = append(devices, &d)
	}

	return devices, nil
}

func (r *deviceRepo) GetDevicesByPrice(_ context.Context, min uint, max uint) ([]*pb.Device, error) {
	query := `SELECT * FROM devices WHERE price BETWEEN $1 AND $2`

	rows, err := r.db.Query(query, min, max)
	if err != nil {
		return nil, err
	}

	var devices []*pb.Device
	for rows.Next() {
		d := pb.Device{}
		if err = rows.Scan(&d.UUID, &d.Title, &d.Description, &d.Price, &d.Manufacturer, &d.Amount); err != nil {
			return nil, err
		}
		devices = append(devices, &d)
	}

	return devices, nil
}

func (r *deviceRepo) GetDeviceByUUIDWithAmount(_ context.Context, deviceUUID string, amount uint32) (*pb.Device, error) {
	query := `UPDATE devices SET amount = amount - $1 WHERE uuid = $2 RETURNING *`

	var device pb.Device

	if err := r.db.QueryRow(query, amount, deviceUUID).Scan(&device.UUID, &device.Title, &device.Description, &device.Price, &device.Manufacturer, &device.Amount); err != nil {
		return nil, err
	}

	return &device, nil
}
