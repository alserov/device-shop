package helpers

import (
	"context"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/order-service/internal/utils"
	"github.com/alserov/device-shop/order-service/pkg/entity"
	pb "github.com/alserov/device-shop/proto/gen"
)

func ChangeBalance(ctx context.Context, chErr chan<- error, userAddr string, order *entity.CreateOrderReqWithDevices) {
	cl, cc, err := client.DialUser(userAddr)
	if err != nil {
		chErr <- err
	}
	defer cc.Close()

	_, err = cl.DebitBalance(ctx, &pb.DebitBalanceReq{
		Cash:     float32(utils.CountOrderPrice(order.Devices)),
		UserUUID: order.UserUUID,
	})
	if err != nil {
		chErr <- err
	}
}
