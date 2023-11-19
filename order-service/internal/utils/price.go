package utils

import (
	pb "github.com/alserov/device-shop/proto/gen"
)

func CountOrderPrice(items []*pb.Device) float32 {
	var price float32
	for _, v := range items {
		price += v.Price * float32(v.Amount)
	}
	return price
}
