package service

import (
	"context"
	brokermock "github.com/alserov/device-shop/order-service/internal/broker/manager/mocks"
	repomock "github.com/alserov/device-shop/order-service/internal/db/mocks"
	repo "github.com/alserov/device-shop/order-service/internal/db/models"
	"github.com/alserov/device-shop/order-service/internal/service/models"
	"github.com/alserov/device-shop/order-service/internal/utils/status"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateOrder(t *testing.T) {
	createOrderReq := models.CreateOrderReq{
		OrderPrice: 100,
		UserUUID:   "uuid",
		OrderDevices: []*models.OrderDevice{
			{
				DeviceUUID: "uuid",
				Amount:     1,
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	broker := brokermock.NewMockTxManager(ctrl)
	broker.EXPECT().CreateOrderTx(gomock.Any())

	s := NewService(nil, broker, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	createdOrder, err := s.CreateOrder(ctx, createOrderReq)
	require.NoError(t, err)

	require.NotEqual(t, "", createdOrder.OrderUUID)
}

func TestCheckOrder(t *testing.T) {
	checkOrderReq := models.CheckOrderReq{
		OrderUUID: "uuid",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()

	store := repomock.NewMockOrderRepo(ctrl)
	store.EXPECT().CheckOrder(gomock.Any(), gomock.Any()).Return(repo.CheckOrderRes{
		Status:     status.PENDING_CODE,
		OrderPrice: 100,
		CreatedAt:  &now,
		OrderDevices: []*repo.OrderDevice{
			{
				DeviceUUID: "uuid",
				Amount:     1,
			},
		},
	}, nil).Times(1)

	s := NewService(store, nil, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	order, err := s.CheckOrder(ctx, checkOrderReq)
	require.NoError(t, err)

	require.Equal(t, now, *order.CreatedAt)
	require.Len(t, order.OrderDevices, 1)
	require.NotEmpty(t, order.Status)
	require.Greater(t, order.OrderPrice, float32(0))
}
