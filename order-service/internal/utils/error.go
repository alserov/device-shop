package utils

import (
	"context"
	"fmt"
	"github.com/alserov/device-shop/order-service/internal/helpers"
	"github.com/alserov/device-shop/order-service/pkg/entity"
	"log"
)

type RequestError struct {
	RequestID uint32
	Err       error
}

func (err *RequestError) Handle(ctx context.Context, order *entity.CreateOrderReqWithDevices, deviceAddr string, userAddr string) error {
	log.Println(err.Err, err.RequestID)
	switch err.RequestID {
	case 1:
		for _, d := range order.Devices {
			if err := helpers.RollBackDeviceAmount(ctx, d.UUID, d.Amount, deviceAddr); err != nil {
				log.Println(fmt.Errorf("failed to rollback device with UUID: %s \t Amount: %d", d.UUID, d.Amount))
			}
		}
		return err.Err
	case 2:
		price := CountOrderPrice(order.Devices)
		if err := helpers.RollBackBalance(ctx, order.UserUUID, price, userAddr); err != nil {
			log.Println(fmt.Errorf("failed to rollback balance with UserUUID: %s \t Cash: %f", order.UserUUID, price))
		}
		return err.Err
	default:
		log.Println("Unexpected error ", err)
		return err.Err
	}
}
