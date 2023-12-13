package service

import (
	"context"

	repomock "github.com/alserov/device-shop/device-service/internal/db/mocks"
	repo "github.com/alserov/device-shop/device-service/internal/db/models"
	"github.com/alserov/device-shop/device-service/internal/service/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateDevice(t *testing.T) {
	createDeviceReq := models.CreateDeviceReq{
		UUID:         "uuid",
		Title:        "title",
		Description:  "desc",
		Price:        100,
		Manufacturer: "manu",
		Amount:       5,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := repomock.NewMockDeviceRepo(ctrl)
	store.EXPECT().CreateDevice(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	s := NewService(store, nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := s.CreateDevice(ctx, createDeviceReq)
	require.NoError(t, err)
}

func TestGetDeviceByUUID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := repomock.NewMockDeviceRepo(ctrl)
	store.EXPECT().GetDeviceByUUID(gomock.Any(), gomock.Any()).Return(repo.Device{
		UUID:         "uuid",
		Title:        "title",
		Description:  "desc",
		Price:        100,
		Manufacturer: "manu",
		Amount:       5,
	}, nil).Times(1)

	s := NewService(store, nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	device, err := s.GetDeviceByUUID(ctx, "uuid")
	require.NoError(t, err)

	require.NotEmpty(t, device.UUID, device.Title, device.Manufacturer, device.Description)
	require.Greater(t, device.Amount, uint32(0))
	require.Greater(t, device.Price, float32(0))
}
