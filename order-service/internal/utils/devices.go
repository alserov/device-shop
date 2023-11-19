package utils

import (
	"context"
	"fmt"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/order-service/pkg/entity"
	pb "github.com/alserov/device-shop/proto/gen"
	"log"
	"sync"
	"time"
)

func FetchDevices(ctx context.Context, deviceAddr string, devices []*pb.OrderDevice, target *entity.CreateOrderReqWithDevices) error {
	cl, cc, err := client.DialDevice(deviceAddr)
	if err != nil {
		return err
	}
	defer cc.Close()

	var (
		wg        = &sync.WaitGroup{}
		chErr     = make(chan error, 1)
		chDevices = make(chan *pb.Device, len(devices))
	)
	wg.Add(len(devices))

	for _, d := range devices {
		d := d
		go func() {
			defer wg.Done()
			device, err := cl.GetDeviceByUUIDWithAmount(ctx, &pb.GetDeviceByUUIDWithAmountReq{
				DeviceUUID: d.DeviceUUID,
				Amount:     d.Amount,
			})
			if err != nil {
				chErr <- err
				return
			}
			chDevices <- device
		}()
	}

	go func() {
		wg.Wait()
		close(chDevices)
		close(chErr)
	}()

	go func() {
		for device := range chDevices {
			target.Devices = append(target.Devices, device)
		}
	}()

	for err = range chErr {
		return err
	}
	return nil
}

func RollbackDeviceAmountPB(devices []*pb.Device, addr string) {
	for _, d := range devices {
		if err := rollbackAmount(d.UUID, d.Amount, addr); err != nil {
			log.Println(fmt.Errorf("failed to rollback device with UUID: %s \t Amount: %d", d.UUID, d.Amount))
		}
	}
}

func RollbackDeviceAmount(devices []*entity.OrderDevice, addr string) {
	for _, d := range devices {
		if err := rollbackAmount(d.DeviceUUID, d.Amount, addr); err != nil {
			log.Println(fmt.Errorf("failed to rollback device with UUID: %s \t Amount: %d", d.DeviceUUID, d.Amount))
		}
	}
}

func rollbackAmount(deviceUUID string, amount uint32, deviceAddr string) error {
	cl, cc, err := client.DialDevice(deviceAddr)
	if err != nil {
		return err
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = cl.IncreaseDeviceAmountByUUID(ctx, &pb.IncreaseDeviceAmountByUUIDReq{
		DeviceUUID: deviceUUID,
		Amount:     amount,
	})
	if err != nil {
		return err
	}

	return nil
}
