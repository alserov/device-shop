package service

import (
	"context"
	"database/sql"
	"github.com/alserov/admin-service/internal/db"
	"github.com/alserov/admin-service/internal/db/postgres"
)

type Service interface {
	CreateDevice(context.Context, CreateDeviceReq) error
	DeleteDevice(context.Context, string) error
	UpdateDevice(context.Context, UpdateDeviceReq) error
}

func NewService(db *sql.DB) Service {
	return &service{
		db: postgres.NewRepo(db),
	}
}

type service struct {
	db db.AdminRepo
}

func (s service) CreateDevice(ctx context.Context, req CreateDeviceReq) error {
	err := s.db.CreateDevice(ctx, db.Device{
		Title:        req.Title,
		Price:        req.Price,
		Manufacturer: req.Manufacturer,
		Description:  req.Description,
		Amount:       req.Amount,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s service) DeleteDevice(ctx context.Context, uuid string) error {
	err := s.db.DeleteDevice(ctx, uuid)
	if err != nil {
		return err
	}

	return nil
}

func (s service) UpdateDevice(ctx context.Context, req UpdateDeviceReq) error {
	err := s.db.UpdateDevice(ctx, db.UpdateDevice{
		UUID:        req.UUID,
		Description: req.Description,
		Title:       req.Title,
		Price:       req.Price,
	})
	if err != nil {
		return err
	}

	return nil
}
