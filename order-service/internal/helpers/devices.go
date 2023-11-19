package helpers

import (
	"context"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/order-service/internal/utils"
	pb "github.com/alserov/device-shop/proto/gen"
	"sync"
)

func FetchDevices(ctx context.Context, chErr chan<- *utils.RequestError, wg *sync.WaitGroup, deviceAddr string, devices []*pb.OrderDevice, target []*pb.Device) {
	defer wg.Done()
	cl, cc, err := client.DialDevice(deviceAddr)
	if err != nil {
		chErr <- &utils.RequestError{
			RequestID: 1,
			Err:       err,
		}
	}
	defer cc.Close()

	localWg := &sync.WaitGroup{}
	localWg.Add(len(devices))

	chDevices := make(chan *pb.Device, len(devices))

	for _, d := range devices {
		d := d
		go func() {
			defer localWg.Done()
			device, err := cl.GetDeviceByUUIDWithAmount(ctx, &pb.GetDeviceByUUIDWithAmountReq{
				DeviceUUID: d.DeviceUUID,
				Amount:     d.Amount,
			})
			if err != nil {
				chErr <- &utils.RequestError{
					RequestID: 1,
					Err:       err,
				}
			}
			chDevices <- device
		}()
	}

	go func() {
		localWg.Wait()
		close(chDevices)
	}()

	for device := range chDevices {
		target = append(target, device)
	}
}

func RollBackDeviceAmount(ctx context.Context, deviceUUID string, amount uint32, deviceAddr string) error {
	cl, cc, err := client.DialDevice(deviceAddr)
	if err != nil {
		return err
	}
	defer cc.Close()

	_, err = cl.IncreaseDeviceAmountByUUID(ctx, &pb.IncreaseDeviceAmountByUUIDReq{
		DeviceUUID: deviceUUID,
		Amount:     amount,
	})
	if err != nil {
		return err
	}

	return nil
}
