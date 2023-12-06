package broker

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/IBM/sarama"

	"github.com/alserov/device-shop/user-service/internal/broker/consumer"
	"github.com/alserov/device-shop/user-service/internal/db"
	"github.com/alserov/device-shop/user-service/internal/service/models"
	"github.com/alserov/device-shop/user-service/internal/utils/converter"

	"log/slog"
)

type TxWorker struct {
	txs  map[string]*sql.Tx
	conv *converter.ServiceConverter
}

func NewTxWorker() *TxWorker {
	return &TxWorker{
		conv: converter.NewServiceConverter(),
	}
}

const (
	succeededStatus = 1
)

// TODO: saga in user service debit balance

func (t *TxWorker) MustStartTxWorker(brokerAddr string, topic string, repo db.UserRepo, log *slog.Logger) {
	cons, err := sarama.NewConsumer([]string{brokerAddr}, &sarama.Config{})
	if err != nil {
		panic("failed to start kafka consumer: " + err.Error())
	}

	msgs, err := consumer.Subscribe(topic, cons)
	if err != nil {
		panic("failed to subscribe on topic: " + err.Error())
	}

	for msg := range msgs {
		var txBalanceReq models.WorkerBalanceReq
		if err = json.Unmarshal(msg, &txBalanceReq); err != nil {
			log.Error("failed to unmarshall balance req: " + err.Error())
			continue
		}

		if _, ok := t.txs[txBalanceReq.TxUUID]; ok {
			switch txBalanceReq.Status {
			case succeededStatus:
				t.txs[txBalanceReq.TxUUID].Commit()
			default:
				delete(t.txs, txBalanceReq.TxUUID)
			}
			continue
		}

		tx, err = repo.DebitBalanceTx(context.Background(), t.conv.Balance.WorkerBalanceReqToRepo(txBalanceReq))
		if err != nil {

		}
	}
}
