package service

import (
	"context"

	repomock "github.com/alserov/device-shop/collection-service/internal/db/mocks"
	"github.com/alserov/device-shop/collection-service/internal/db/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := repomock.NewMockCollectionRepo(ctrl)
	store.EXPECT().GetCart(gomock.Any(), gomock.Any()).Return([]*models.Device{
		{
			UUID:         "uuid",
			Title:        "title",
			Description:  "desc",
			Price:        100,
			Manufacturer: "manu",
			Amount:       5,
		},
	}, nil).Times(1)

	s := NewService(store, nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	cart, err := s.GetCart(ctx, "uuid")
	require.NoError(t, err)

	require.Len(t, cart, 1)

	device := cart[0]
	require.NotEmpty(t, device.UUID, device.Title, device.Description, device.Manufacturer)
	require.Greater(t, device.Amount, uint32(0))
	require.Greater(t, device.Price, float32(0))
}
