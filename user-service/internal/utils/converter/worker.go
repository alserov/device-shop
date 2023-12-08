package converter

import (
	"github.com/alserov/device-shop/user-service/internal/broker/worker/models"
	repo "github.com/alserov/device-shop/user-service/internal/db/models"
)

type BrokerConverter struct {
}

func NewBrokerConverter() *BrokerConverter {
	return &BrokerConverter{}
}

func (s *BrokerConverter) WorkerBalanceReqToRepo(req models.Request) repo.BalanceReq {
	return repo.BalanceReq{
		UserUUID: req.UserUUID,
		Cash:     req.OrderPrice,
	}
}
