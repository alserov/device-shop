package service

import (
	"context"
	"errors"

	brokermock "github.com/alserov/device-shop/order-service/internal/broker/manager/mocks"
	"github.com/alserov/device-shop/order-service/internal/db"
	repomock "github.com/alserov/device-shop/order-service/internal/db/mocks"
	repo "github.com/alserov/device-shop/order-service/internal/db/models"
	"github.com/alserov/device-shop/order-service/internal/service/models"
	"github.com/alserov/device-shop/order-service/internal/utils/status"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"log/slog"
	"testing"
	"time"
)

func TestCreateOrder(t *testing.T) {
	createOrderReq := models.CreateOrderReq{
		OrderPrice: 100,
		UserUUID:   "uuid",
		OrderDevices: []models.OrderDevice{
			{
				DeviceUUID: "uuid",
				Amount:     1,
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	broker := brokermock.NewMockTxManager(ctrl)
	broker.EXPECT().CreateOrderTx(gomock.Any()).Return(nil).Times(1)

	s := NewService(nil, broker, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	orderUUID, err := s.CreateOrder(ctx, createOrderReq)
	require.NoError(t, err)

	require.NotEmpty(t, orderUUID)
}

func TestCancelOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := repomock.NewMockOrderRepo(ctrl)
	storeTx := repomock.NewMockSqlTx(ctrl)
	storeTx.
		EXPECT().
		Commit().
		Return(nil).
		Times(2)
	store.
		EXPECT().
		CancelOrderTx(gomock.Any(), gomock.Any()).
		Return(&db.CancelOrder{Price: 100, UserUUID: "uuid", Tx: storeTx}, nil).
		Times(1)
	store.EXPECT().
		CancelOrderDevicesTx(gomock.Any(), gomock.Any()).
		Return(&db.CancelOrderDevices{
			Devices: []repo.OrderDevice{{DeviceUUID: "uuid", Amount: 1}}, Tx: storeTx}, nil).
		Times(1)

	manager := brokermock.NewMockTxManager(ctrl)
	manager.
		EXPECT().
		CancelOrderTx(gomock.Any()).
		Return(nil).
		Times(1)

	s := NewService(store, manager, nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := s.CancelOrder(ctx, "uuid")
	require.NoError(t, err)
}

func TestCancelOrder_with_broker_error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := repomock.NewMockOrderRepo(ctrl)
	manager := brokermock.NewMockTxManager(ctrl)
	s := NewService(store, manager, &slog.Logger{})

	storeTx := repomock.NewMockSqlTx(ctrl)
	storeTx.
		EXPECT().
		Rollback().
		Return(nil).
		Times(2)
	store.
		EXPECT().
		CancelOrderTx(gomock.Any(), gomock.Any()).
		Return(&db.CancelOrder{Price: 100, UserUUID: "uuid", Tx: storeTx}, nil).
		Times(1)
	store.EXPECT().
		CancelOrderDevicesTx(gomock.Any(), gomock.Any()).
		Return(&db.CancelOrderDevices{
			Devices: []repo.OrderDevice{{DeviceUUID: "uuid", Amount: 1}}, Tx: storeTx}, nil).
		Times(1)

	manager.
		EXPECT().
		CancelOrderTx(gomock.Any()).
		Return(errors.New("error")).
		Times(1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := s.CancelOrder(ctx, "uuid")
	require.Error(t, err)
}

func TestCancelOrder_with_repo_error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := repomock.NewMockOrderRepo(ctrl)
	s := NewService(store, nil, &slog.Logger{})

	// cancelOrderTx ERROR
	storeTx := repomock.NewMockSqlTx(ctrl)
	storeTx.
		EXPECT().
		Rollback().
		Return(nil).
		Times(2)
	store.
		EXPECT().
		CancelOrderTx(gomock.Any(), gomock.Any()).
		Return(&db.CancelOrder{Tx: storeTx}, errors.New("error")).
		Times(1)
	store.EXPECT().
		CancelOrderDevicesTx(gomock.Any(), gomock.Any()).
		Return(&db.CancelOrderDevices{
			Devices: []repo.OrderDevice{{DeviceUUID: "uuid", Amount: 1}}, Tx: storeTx}, nil).
		Times(1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := s.CancelOrder(ctx, "uuid")
	require.Error(t, err)

	// CancelOrderDevicesTx ERROR
	storeTx.
		EXPECT().
		Rollback().
		Return(nil).
		Times(2)
	store.
		EXPECT().
		CancelOrderTx(gomock.Any(), gomock.Any()).
		Return(&db.CancelOrder{Price: 100, UserUUID: "uuid", Tx: storeTx}, nil).
		Times(1)
	store.EXPECT().
		CancelOrderDevicesTx(gomock.Any(), gomock.Any()).
		Return(&db.CancelOrderDevices{Devices: nil, Tx: storeTx}, errors.New("error")).
		Times(1)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err = s.CancelOrder(ctx, "uuid")
	require.Error(t, err)
}

func TestUpdateOrder(t *testing.T) {
	updateOrderReq := models.UpdateOrderReq{
		OrderUUID: "uuid",
		Status:    status.DELIVERING,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := repomock.NewMockOrderRepo(ctrl)
	store.EXPECT().UpdateOrder(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

	s := NewService(store, nil, nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := s.UpdateOrder(ctx, updateOrderReq)
	require.NoError(t, err)
}

func TestCheckOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()

	store := repomock.NewMockOrderRepo(ctrl)
	store.EXPECT().CheckOrder(gomock.Any(), gomock.Any()).Return(repo.CheckOrderRes{
		Status:     status.PENDING_CODE,
		OrderPrice: 100,
		CreatedAt:  &now,
		OrderDevices: []repo.OrderDevice{
			{
				DeviceUUID: "uuid",
				Amount:     1,
			},
		},
	}, nil).Times(1)

	s := NewService(store, nil, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	order, err := s.CheckOrder(ctx, "uuid")
	require.NoError(t, err)

	require.Equal(t, now, *order.CreatedAt)
	require.Len(t, order.OrderDevices, 1)
	require.NotEmpty(t, order.Status)
	require.Greater(t, order.OrderPrice, float32(0))
}
