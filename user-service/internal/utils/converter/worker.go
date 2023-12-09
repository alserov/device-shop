package converter

import (
	"github.com/alserov/device-shop/user-service/internal/broker/worker/models"
	repo "github.com/alserov/device-shop/user-service/internal/db/models"
)

type brokerConverter struct {
}

type BrokerConverter interface {
	WorkerBalanceReqToRepo(req models.Request) repo.BalanceReq
}

func NewBrokerConverter() BrokerConverter {
	return &brokerConverter{}
}

func (s *brokerConverter) WorkerBalanceReqToRepo(req models.Request) repo.BalanceReq {
	return repo.BalanceReq{
		UserUUID: req.UserUUID,
		Cash:     req.OrderPrice,
	}
}
