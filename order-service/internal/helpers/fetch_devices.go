package helpers

import (
	"context"
	"github.com/alserov/device-shop/gateway/pkg/client"
	pb "github.com/alserov/device-shop/proto/gen"
	"sync"
)

func FetchDevices(ctx context.Context, chDevices chan<- *pb.Device, chErr chan<- error, wg *sync.WaitGroup, deviceAddr string, deviceUUIDs []string) {
	cl, cc, err := client.DialDevice(deviceAddr)
	if err != nil {
		chErr <- err
	}
	defer cc.Close()

	for _, v := range deviceUUIDs {
		v := v
		go func() {
			defer wg.Done()
			device, err := cl.GetDeviceByUUID(ctx, &pb.UUIDReq{
				UUID: v,
			})
			if err != nil {
				chErr <- err
			}
			chDevices <- device
		}()
	}

	go func() {
		wg.Wait()
		close(chDevices)
	}()
}
