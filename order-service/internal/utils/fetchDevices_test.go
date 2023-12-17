package utils

import (
	"context"
	"github.com/alserov/device-shop/proto/gen/order"

	deviceservicemock "github.com/alserov/device-shop/order-service/internal/service/mocks"
	"github.com/alserov/device-shop/order-service/internal/service/models"
	"github.com/alserov/device-shop/proto/gen/device"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestFetchDevicesFromOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := deviceservicemock.NewMockDevicesClient(ctrl)
	c.
		EXPECT().
		GetDeviceByUUID(gomock.Any(), gomock.Eq(&device.GetDeviceByUUIDReq{UUID: "uuid"})).
		Return(&device.Device{
			UUID:         "uuid",
			Title:        "title",
			Description:  "desc",
			Price:        100,
			Manufacturer: "manu",
			Amount:       5,
		}, nil).
		Times(3)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	fetchedDevices, err := FetchDevicesFromOrder(ctx, c, []models.OrderDevice{{
		DeviceUUID: "uuid",
		Amount:     5,
	}, {
		DeviceUUID: "uuid",
		Amount:     5,
	}, {
		DeviceUUID: "uuid",
		Amount:     5,
	}})
	require.NoError(t, err)
	require.Len(t, fetchedDevices, 3)

	for _, d := range fetchedDevices {
		require.NotEmpty(t, d.UUID)
		require.NotEmpty(t, d.UUID)
		require.NotEmpty(t, d.Manufacturer)
		require.NotEmpty(t, d.Title)
		require.NotEmpty(t, d.Description)
		require.Greater(t, d.Amount, uint32(0))
		require.Greater(t, d.Price, float32(0))
	}
}

func BenchmarkFetchDevicesFromOrder(b *testing.B) {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	c := deviceservicemock.NewMockDevicesClient(ctrl)
	c.
		EXPECT().
		GetDeviceByUUID(gomock.Any(), gomock.Eq(&device.GetDeviceByUUIDReq{UUID: "uuid"})).
		Return(&device.Device{
			UUID:         "uuid",
			Title:        "title",
			Description:  "desc",
			Price:        100,
			Manufacturer: "manu",
			Amount:       5,
		}, nil).
		Times(b.N)

	for i := 0; i < b.N; i++ {
		FetchDevicesFromOrder(context.Background(), c, []models.OrderDevice{{
			DeviceUUID: "uuid",
			Amount:     5,
		}})
	}
}

func TestCountOrderPrice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := deviceservicemock.NewMockDevicesClient(ctrl)
	c.
		EXPECT().
		GetDeviceByUUID(gomock.Any(), gomock.Eq(&device.GetDeviceByUUIDReq{UUID: "uuid"})).
		Return(&device.Device{
			UUID:         "uuid",
			Title:        "title",
			Description:  "desc",
			Price:        100,
			Manufacturer: "manu",
			Amount:       5,
		}, nil).
		Times(3)

	price, err := CountOrderPrice(context.Background(), c, []*order.OrderDevice{
		{
			DeviceUUID: "uuid",
			Amount:     5,
		},
		{
			DeviceUUID: "uuid",
			Amount:     5,
		},
		{
			DeviceUUID: "uuid",
			Amount:     5,
		},
	})
	require.NoError(t, err)
	require.Equal(t, float32(1500), price)
}

func BenchmarkCountOrderPrice(b *testing.B) {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	c := deviceservicemock.NewMockDevicesClient(ctrl)
	c.
		EXPECT().
		GetDeviceByUUID(gomock.Any(), gomock.Eq(&device.GetDeviceByUUIDReq{UUID: "uuid"})).
		Return(&device.Device{
			UUID:         "uuid",
			Title:        "title",
			Description:  "desc",
			Price:        100,
			Manufacturer: "manu",
			Amount:       5,
		}, nil).
		Times(b.N)

	for i := 0; i < b.N; i++ {
		CountOrderPrice(context.Background(), c, []*order.OrderDevice{
			{
				DeviceUUID: "uuid",
				Amount:     5,
			},
		})
	}
}
