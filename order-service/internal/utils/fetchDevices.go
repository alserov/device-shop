package utils

import (
	"context"
	"github.com/alserov/device-shop/order-service/internal/service/models"
	"github.com/alserov/device-shop/proto/gen/device"
	"github.com/alserov/device-shop/proto/gen/order"
	"sync"
)

func FetchDevicesFromOrder(ctx context.Context, cl device.DevicesClient, devicesFromOrder []models.OrderDevice) ([]*device.Device, error) {
	var (
		wg      = &sync.WaitGroup{}
		chErr   = make(chan error)
		devices = make([]*device.Device, 0, len(devicesFromOrder))
	)

	wg.Add(len(devicesFromOrder))

	for _, d := range devicesFromOrder {
		d := d
		go func() {
			defer wg.Done()
			device, err := cl.GetDeviceByUUID(ctx, &device.GetDeviceByUUIDReq{
				UUID: d.DeviceUUID,
			})
			if err != nil {
				chErr <- err
			}
			device.Amount = d.Amount
			devices = append(devices, device)
		}()
	}

	go func() {
		wg.Wait()
		close(chErr)
	}()

	for err := range chErr {
		return nil, err
	}

	return devices, nil
}

func CountOrderPrice(ctx context.Context, cl device.DevicesClient, orderDevices []*order.OrderDevice) (float32, error) {
	var (
		price float32
		wg    = &sync.WaitGroup{}
		mu    = &sync.Mutex{}
		chErr = make(chan error)
	)

	wg.Add(len(orderDevices))

	for _, od := range orderDevices {
		od := od
		go func() {
			defer wg.Done()
			d, err := cl.GetDeviceByUUID(ctx, &device.GetDeviceByUUIDReq{
				UUID: od.DeviceUUID,
			})
			if err != nil {
				chErr <- err
				return
			}
			mu.Lock()
			defer mu.Unlock()
			price += d.Price * float32(od.Amount)
		}()
	}

	go func() {
		wg.Wait()
		close(chErr)
	}()

	for err := range chErr {
		return 0, err
	}

	return price, nil
}
