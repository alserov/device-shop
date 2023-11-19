package helpers

import (
	"context"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/order-service/internal/utils"
	"github.com/alserov/device-shop/order-service/pkg/entity"
	pb "github.com/alserov/device-shop/proto/gen"
	"sync"
)

func ChangeBalance(ctx context.Context, chErr chan<- *entity.RequestError, wg *sync.WaitGroup, userAddr string, order *entity.CreateOrderReqWithDevices) {
	defer wg.Done()
	cl, cc, err := client.DialUser(userAddr)
	if err != nil {
		chErr <- &entity.RequestError{
			RequestID: 2,
			Err:       err,
		}
	}
	defer cc.Close()

	_, err = cl.DebitBalance(ctx, &pb.DebitBalanceReq{
		Cash:     utils.CountOrderPrice(order.Devices),
		UserUUID: order.UserUUID,
	})
	if err != nil {
		chErr <- &entity.RequestError{
			RequestID: 2,
			Err:       err,
		}
	}
}
