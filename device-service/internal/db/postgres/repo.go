package postgres

import (
	"context"
	"database/sql"
	"errors"
	pb "github.com/alserov/device-shop/proto/gen"
)

type Repository interface {
	CreateDevice(context.Context, *pb.Device) error
	DeleteDevice(context.Context, string) error
	UpdateDevice(context.Context, *pb.UpdateDeviceReq) error
	GetAllDevices(context.Context, uint32, uint32) ([]*pb.Device, error)
	GetDevicesByTitle(context.Context, string) ([]*pb.Device, error)
	GetDeviceByUUID(context.Context, string) (*pb.Device, error)
	GetDevicesByManufacturer(context.Context, string) ([]*pb.Device, error)
	GetDevicesByPrice(context.Context, uint, uint) ([]*pb.Device, error)
	GetDeviceByUUIDWithAmount(context.Context, string, uint32) (*pb.Device, error)
	IncreaseDeviceAmountByUUID(ctx context.Context, deviceUUID string, amount uint32) error
}

type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) CreateDevice(ctx context.Context, device *pb.Device) error {
	query := `INSERT INTO devices (uuid, title, description, price, manufacturer, amount) VALUES($1,$2,$3,$4,$5,$6)`

	_, err := r.db.Exec(query, device.UUID, device.Title, device.Description, device.Price, device.Manufacturer, device.Amount)
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

func (r *repo) UpdateDevice(ctx context.Context, device *pb.UpdateDeviceReq) error {
	query := `UPDATE devices SET title = $1, description = $2, price = $3 WHERE uuid = $4`

	_, err := r.db.Exec(query, device.Title, device.Description, device.Price, device.UUID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetAllDevices(ctx context.Context, index uint32, amount uint32) ([]*pb.Device, error) {
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

func (r *repo) GetDevicesByTitle(ctx context.Context, title string) ([]*pb.Device, error) {
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

func (r *repo) GetDeviceByUUID(ctx context.Context, uuid string) (*pb.Device, error) {
	query := `SELECT * FROM devices WHERE uuid = $1`

	d := pb.Device{}

	err := r.db.QueryRow(query, uuid).Scan(&d.UUID, &d.Title, &d.Description, &d.Price, &d.Manufacturer, &d.Amount)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return &d, nil
}

func (r *repo) GetDevicesByManufacturer(ctx context.Context, manu string) ([]*pb.Device, error) {
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

func (r *repo) GetDevicesByPrice(ctx context.Context, min uint, max uint) ([]*pb.Device, error) {
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

func (r *repo) GetDeviceByUUIDWithAmount(ctx context.Context, deviceUUID string, amount uint32) (*pb.Device, error) {
	query := `UPDATE devices SET amount = amount - $1 WHERE uuid = $2 RETURNING *`

	var device pb.Device

	if err := r.db.QueryRow(query, amount, deviceUUID).Scan(&device.UUID, &device.Title, &device.Description, &device.Price, &device.Manufacturer, &device.Amount); err != nil {
		return nil, err
	}

	return &device, nil
}

func (r *repo) IncreaseDeviceAmountByUUID(ctx context.Context, deviceUUID string, amount uint32) error {
	query := `UPDATE devices SET amount = amount + $1 WHERE uuid = $2`

	_, err := r.db.Exec(query, amount, deviceUUID)
	if err != nil {
		return err
	}

	return nil
}
