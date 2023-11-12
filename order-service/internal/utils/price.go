package utils

import (
	"github.com/alserov/shop/gateway/pkg/models"
	"sync"
	"sync/atomic"
)

func CountOrderPrice(items []*models.Device) uint {
	var price uint32
	wg := &sync.WaitGroup{}
	wg.Add(len(items))
	for _, v := range items {
		v := v
		go func() {
			atomic.AddUint32(&price, v.Price)
			wg.Done()
		}()
	}
	wg.Wait()
	return uint(price)
}
