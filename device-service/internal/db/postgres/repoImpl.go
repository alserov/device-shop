package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alserov/device-shop/device-service/internal/db"
	"github.com/alserov/device-shop/device-service/internal/db/models"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"sync"
)

func NewRepo(db *sql.DB, log *slog.Logger) db.DeviceRepo {
	return &repo{
		db:  db,
		log: log,
	}
}

type repo struct {
	log *slog.Logger
	db  *sql.DB
}

const (
	internalError = "internal error"
	notFound      = "device found"

	invalidDeviceAmountError  = "invalid device amount"
	amountConstraintErrorCode = "23514"
)

func (r *repo) GetAllDevices(_ context.Context, index uint32, amount uint32) ([]*models.Device, error) {
	query := `SELECT * FROM devices LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, amount, index)
	if errors.Is(sql.ErrNoRows, err) {
		return nil, status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		return nil, err
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
	if errors.Is(sql.ErrNoRows, err) {
		return nil, status.Error(codes.NotFound, notFound)
	}
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
	if errors.Is(sql.ErrNoRows, err) {
		return models.Device{}, status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		return models.Device{}, err
	}

	return d, nil
}

func (r *repo) GetDevicesByManufacturer(_ context.Context, manu string) ([]*models.Device, error) {
	query := `SELECT * FROM devices WHERE manufacturer = $1`

	rows, err := r.db.Query(query, manu)
	if errors.Is(sql.ErrNoRows, err) {
		return nil, status.Error(codes.NotFound, notFound)
	}
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
	if errors.Is(sql.ErrNoRows, err) {
		return nil, status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		return nil, status.Error(codes.Internal, internalError)
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
	err := r.db.QueryRow(query, amount, deviceUUID).
		Scan(&device.UUID, &device.Title, &device.Description, &device.Price, &device.Manufacturer, &device.Amount)
	if errors.Is(sql.ErrNoRows, err) {
		return nil, status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		return nil, status.Error(codes.Internal, internalError)
	}

	return &device, nil
}

func (r *repo) CreateDevice(_ context.Context, device models.Device) error {
	op := "repo.CreateDevice"

	query := `INSERT INTO devices (uuid, title, description, price, manufacturer, amount) VALUES($1,$2,$3,$4,$5,$6)`

	_, err := r.db.Exec(query, device.UUID, device.Title, device.Description, device.Price, device.Manufacturer, device.Amount)
	if errors.Is(sql.ErrNoRows, err) {
		return status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed to create device", slog.String("error", err.Error()), slog.String("op", op))
		return status.Error(codes.Internal, internalError)
	}

	return nil
}

func (r *repo) DeleteDevice(_ context.Context, uuid string) error {
	op := "repo.DeleteDevice"

	query := `DELETE FROM devices WHERE uuid = $1`

	_, err := r.db.Exec(query, uuid)
	if errors.Is(sql.ErrNoRows, err) {
		return status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed to delete device", slog.String("error", err.Error()), slog.String("op", op))
		return status.Error(codes.Internal, internalError)
	}

	return nil
}

func (r *repo) UpdateDevice(_ context.Context, device models.UpdateDevice) error {
	op := "repo.UpdateDevice"

	query := `UPDATE devices SET title = $1, description = $2, price = $3 WHERE uuid = $4`

	_, err := r.db.Exec(query, device.Title, device.Description, device.Price, device.UUID)
	if errors.Is(sql.ErrNoRows, err) {
		return status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed to update device", slog.String("error", err.Error()), slog.String("op", op))
		return status.Error(codes.Internal, internalError)
	}

	return nil
}

func (r *repo) IncreaseDeviceAmountByUUID(_ context.Context, deviceUUID string, amount uint32) error {
	op := "repo.IncreaseDeviceAmountByUUID"

	query := `UPDATE devices SET amount = amount + $1 WHERE uuid = $2`

	_, err := r.db.Exec(query, amount, deviceUUID)
	if errors.Is(sql.ErrNoRows, err) {
		return status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed to increase device amount by its uuid", slog.String("error", err.Error()), slog.String("op", op))
		return status.Error(codes.Internal, internalError)
	}

	return nil
}

func (r *repo) DecreaseDevicesAmountTx(ctx context.Context, devices []*models.OrderDevice) (*sql.Tx, error) {
	query := `UPDATE devices SET amount = amount - $1 WHERE uuid = $2`

	tx, err := r.db.Begin()
	if err != nil {
		return tx, status.Error(codes.Internal, internalError)
	}

	var (
		wg    = &sync.WaitGroup{}
		chErr = make(chan error)
	)
	wg.Add(len(devices))

	for _, d := range devices {
		go func(d *models.OrderDevice) {
			defer wg.Done()
			if _, err := tx.Exec(query, d.Amount, d.DeviceUUID); err != nil {
				chErr <- err
			}
		}(d)
	}

	go func() {
		wg.Wait()
		close(chErr)
	}()

	for err = range chErr {
		if errors.Is(sql.ErrNoRows, err) {
			return tx, status.Error(codes.NotFound, notFound)
		}
		switch err.(*pq.Error).Code {
		case amountConstraintErrorCode:
			return tx, status.Error(codes.Canceled, invalidDeviceAmountError)
		default:
			return tx, status.Error(codes.Internal, internalError)
		}
	}

	return tx, nil
}
