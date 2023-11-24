package postgres

import (
	"context"
	"database/sql"
	"github.com/alserov/device-shop/device-service/internal/db"
	pb "github.com/alserov/device-shop/proto/gen"
)

func NewAdminRepo(db *sql.DB) db.AdminRepo {
	return &adminRepo{
		db: db,
	}
}

type adminRepo struct {
	db *sql.DB
}

func (r *adminRepo) CreateDevice(_ context.Context, device *pb.Device) error {
	query := `INSERT INTO devices (uuid, title, description, price, manufacturer, amount) VALUES($1,$2,$3,$4,$5,$6)`

	_, err := r.db.Exec(query, device.UUID, device.Title, device.Description, device.Price, device.Manufacturer, device.Amount)
	if err != nil {
		return err
	}

	return nil
}

func (r *adminRepo) DeleteDevice(_ context.Context, uuid string) error {
	query := `DELETE FROM devices WHERE uuid = $1`

	_, err := r.db.Exec(query, uuid)
	if err != nil {
		return err
	}

	return nil
}

func (r *adminRepo) UpdateDevice(_ context.Context, device *pb.UpdateDeviceReq) error {
	query := `UPDATE devices SET title = $1, description = $2, price = $3 WHERE uuid = $4`

	_, err := r.db.Exec(query, device.Title, device.Description, device.Price, device.UUID)
	if err != nil {
		return err
	}

	return nil
}

func (r *adminRepo) IncreaseDeviceAmountByUUID(_ context.Context, deviceUUID string, amount uint32) error {
	query := `UPDATE devices SET amount = amount + $1 WHERE uuid = $2`

	_, err := r.db.Exec(query, amount, deviceUUID)
	if err != nil {
		return err
	}

	return nil
}
