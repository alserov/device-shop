package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alserov/device-shop/device-service/internal/db"
	"github.com/alserov/device-shop/device-service/internal/db/models"
)

func NewRepo(db *sql.DB) db.DeviceRepo {
	return &repo{
		db: db,
	}
}

type repo struct {
	db *sql.DB
}

func (r *repo) GetAllDevices(_ context.Context, index uint32, amount uint32) ([]*models.Device, error) {
	query := `SELECT * FROM devices LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, amount, index)
	if err != nil {
		return nil, nil
	}

	devices := make([]*models.Device, 0, amount)
	for rows.Next() {
		d := models.Device{}
		if err = rows.Scan(&d.UUID, &d.Title, &d.Description, &d.Price, &d.Manufacturer, &d.Amount); err != nil {
			return nil, err
		}
		devices = append(devices, &d)
	}

	return devices, nil
}

func (r *repo) GetDevicesByTitle(_ context.Context, title string) ([]*models.Device, error) {
	query := `SELECT * FROM devices WHERE title LIKE $1`

	rows, err := r.db.Query(query, "%"+title+"%")
	if err != nil {
		return nil, err
	}

	var devices []*models.Device
	for rows.Next() {
		d := models.Device{}
		if err = rows.Scan(&d.UUID, &d.Title, &d.Description, &d.Price, &d.Manufacturer, &d.Amount); err != nil {
			return nil, err
		}
		devices = append(devices, &d)
	}

	return devices, nil
}

func (r *repo) GetDeviceByUUID(_ context.Context, uuid string) (models.Device, error) {
	query := `SELECT * FROM devices WHERE uuid = $1`

	d := models.Device{}

	err := r.db.QueryRow(query, uuid).Scan(&d.UUID, &d.Title, &d.Description, &d.Price, &d.Manufacturer, &d.Amount)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return models.Device{}, err
	}

	return d, nil
}

func (r *repo) GetDevicesByManufacturer(_ context.Context, manu string) ([]*models.Device, error) {
	query := `SELECT * FROM devices WHERE manufacturer = $1`

	rows, err := r.db.Query(query, manu)
	if err != nil {
		return nil, err
	}

	var devices []*models.Device
	for rows.Next() {
		d := models.Device{}
		if err = rows.Scan(&d.UUID, &d.Title, &d.Description, &d.Price, &d.Manufacturer, &d.Amount); err != nil {
			return nil, err
		}
		devices = append(devices, &d)
	}

	return devices, nil
}

func (r *repo) GetDevicesByPrice(_ context.Context, min uint, max uint) ([]*models.Device, error) {
	query := `SELECT * FROM devices WHERE price BETWEEN $1 AND $2`

	rows, err := r.db.Query(query, min, max)
	if err != nil {
		return nil, err
	}

	var devices []*models.Device
	for rows.Next() {
		d := models.Device{}
		if err = rows.Scan(&d.UUID, &d.Title, &d.Description, &d.Price, &d.Manufacturer, &d.Amount); err != nil {
			return nil, err
		}
		devices = append(devices, &d)
	}

	return devices, nil
}

func (r *repo) GetDeviceByUUIDWithAmount(_ context.Context, deviceUUID string, amount uint32) (*models.Device, error) {
	query := `UPDATE devices SET amount = amount - $1 WHERE uuid = $2 RETURNING *`

	var device models.Device

	if err := r.db.QueryRow(query, amount, deviceUUID).Scan(&device.UUID, &device.Title, &device.Description, &device.Price, &device.Manufacturer, &device.Amount); err != nil {
		return nil, err
	}

	return &device, nil
}

func (r *repo) CreateDevice(_ context.Context, device models.Device) error {
	query := `INSERT INTO devices (uuid, title, description, price, manufacturer, amount) VALUES($1,$2,$3,$4,$5,$6)`

	_, err := r.db.Exec(query, device.UUID, device.Title, device.Description, device.Price, device.Manufacturer, device.Amount)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) DeleteDevice(_ context.Context, uuid string) error {
	query := `DELETE FROM devices WHERE uuid = $1`

	_, err := r.db.Exec(query, uuid)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) UpdateDevice(_ context.Context, device models.UpdateDevice) error {
	query := `UPDATE devices SET title = $1, description = $2, price = $3 WHERE uuid = $4`

	_, err := r.db.Exec(query, device.Title, device.Description, device.Price, device.UUID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) IncreaseDeviceAmountByUUID(_ context.Context, deviceUUID string, amount uint32) error {
	query := `UPDATE devices SET amount = amount + $1 WHERE uuid = $2`

	_, err := r.db.Exec(query, amount, deviceUUID)
	if err != nil {
		return err
	}

	return nil
}