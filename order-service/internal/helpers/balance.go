package helpers

import (
	"context"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/order-service/internal/utils"
	"github.com/alserov/device-shop/order-service/pkg/entity"
	pb "github.com/alserov/device-shop/proto/gen"
	"sync"
)

func ChangeBalance(ctx context.Context, chErr chan<- *utils.RequestError, wg *sync.WaitGroup, userAddr string, order *entity.CreateOrderReqWithDevices) {
	defer wg.Done()
	cl, cc, err := client.DialUser(userAddr)
	if err != nil {
		chErr <- &utils.RequestError{
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
		chErr <- &utils.RequestError{
			RequestID: 2,
			Err:       err,
		}
	}
}

// TODO: finish rollback balance
func RollBackBalance(ctx context.Context, userUUID string, cash float32, userAddr string) error {

}
