package utils

import (
	"context"
	"fmt"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/order-service/pkg/entity"
	pb "github.com/alserov/device-shop/proto/gen"
	"log"
	"time"
)

func ChangeBalance(ctx context.Context, userAddr string, order *entity.CreateOrderReqWithDevices) error {
	cl, cc, err := client.DialUser(userAddr)
	if err != nil {
		return err
	}
	defer cc.Close()

	_, err = cl.DebitBalance(ctx, &pb.DebitBalanceReq{
		Cash:     CountOrderPrice(order.Devices),
		UserUUID: order.UserUUID,
	})
	if err != nil {
		return err
	}

	return nil
}

func RollBackBalance(userUUID string, cash float32, addr string) {
	cl, cc, err := client.DialUser(addr)
	if err != nil {
		log.Println(err)
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = cl.TopUpBalance(ctx, &pb.TopUpBalanceReq{
		Cash:     cash,
		UserUUID: userUUID,
	})
	if err != nil {
		log.Println(fmt.Errorf("failed to rollback balance with UserUUID: %s \t Amount: %f", userUUID, cash))
	}
}
