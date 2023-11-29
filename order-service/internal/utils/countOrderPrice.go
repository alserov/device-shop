package utils

import (
	"context"
	"github.com/alserov/device-shop/gateway/pkg/client"
	pb "github.com/alserov/device-shop/proto/gen"
	"os"
	"sync"
)

func CountPrice(ctx context.Context, devices []*pb.OrderDevice) (float32, error) {
	var (
		price float32
		wg    = &sync.WaitGroup{}
		mu    = &sync.Mutex{}
		chErr = make(chan error)
	)

	cl, cc, err := client.DialDevice(os.Getenv("DEVICE_ADDR"))
	if err != nil {
		return 0, err
	}
	defer cc.Close()

	wg.Add(len(devices))

	for _, od := range devices {
		od := od
		go func() {
			defer wg.Done()
			d, err := cl.GetDeviceByUUID(ctx, &pb.GetDeviceByUUIDReq{
				UUID: od.DeviceUUID,
			})
			if err != nil {
				chErr <- err
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

	for err = range chErr {
		return 0, err
	}

	return price, nil
}
