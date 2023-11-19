package helpers

import (
	"context"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/order-service/pkg/entity"
	pb "github.com/alserov/device-shop/proto/gen"
	"sync"
)

func FetchDevices(ctx context.Context, chErr chan<- *entity.RequestError, wg *sync.WaitGroup, deviceAddr string, devices []*pb.OrderDevice, target []*pb.Device) {
	defer wg.Done()
	cl, cc, err := client.DialDevice(deviceAddr)
	if err != nil {
		chErr <- &entity.RequestError{
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
				chErr <- &entity.RequestError{
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
