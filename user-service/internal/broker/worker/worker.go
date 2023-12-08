package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/user-service/internal/broker"
	"github.com/alserov/device-shop/user-service/internal/broker/producer"
	"github.com/alserov/device-shop/user-service/internal/broker/worker/models"
	"github.com/alserov/device-shop/user-service/internal/db/postgres"
	"google.golang.org/grpc/status"
	"log"

	"github.com/alserov/device-shop/user-service/internal/broker/consumer"
	"github.com/alserov/device-shop/user-service/internal/db"
	"github.com/alserov/device-shop/user-service/internal/utils/converter"

	"log/slog"
)

type TxWorker struct {
	log *slog.Logger

	txs map[string]*sql.Tx

	topicIn  string
	topicOut string

	repo db.UserRepo
	conv *converter.BrokerConverter

	c sarama.Consumer
	p sarama.SyncProducer
}

const (
	internalError = "internal error"
	kafkaClientID = "USER_WORKER"
)

func NewTxWorker(broker *broker.Broker, db *sql.DB, log *slog.Logger) *TxWorker {
	cons, err := sarama.NewConsumer([]string{broker.BrokerAddr}, nil)
	if err != nil {
		panic("failed to start kafka consumer: " + err.Error())
	}

	prod, err := producer.NewProducer([]string{broker.BrokerAddr}, kafkaClientID)
	if err != nil {
		panic("failed to start kafka producer: " + err.Error())
	}

	return &TxWorker{
		log:      log,
		c:        cons,
		p:        prod,
		conv:     converter.NewBrokerConverter(),
		repo:     postgres.NewRepo(db, log),
		topicIn:  broker.Topics.Manager.In,
		topicOut: broker.Topics.Manager.Out,
		txs:      make(map[string]*sql.Tx),
	}
}

const (
	serverFailureStatus = 0
	userFailureStatus   = 1
	successStatus       = 2
)

type txResponse struct {
	// 0 - failed
	// 1 - success
	Status  uint32 `json:"status"`
	Message string `json:"message"`
	Uuid    string `json:"uuid"`
}

func (w *TxWorker) MustStart() {
	msgs, err := consumer.Subscribe(w.topicIn, w.c)
	if err != nil {
		panic("failed to subscribe on topic: " + err.Error())
	}

	w.log.Info("worker is running")
	for msg := range msgs {
		var req models.Request
		if err = json.Unmarshal(msg, &req); err != nil {
			w.log.Error("failed to unmarshall balance req: " + err.Error())
			continue
		}

		fmt.Println(req)

		if _, ok := w.txs[req.TxUUID]; ok {
			switch req.Status {
			case successStatus:
				if err := w.txs[req.TxUUID].Commit(); err != nil {
					log.Println(err)
				}
			default:
				w.txs[req.TxUUID].Rollback()
			}
			delete(w.txs, req.TxUUID)
		} else {
			tx, err := w.repo.DebitBalanceTx(context.Background(), w.conv.WorkerBalanceReqToRepo(req))
			if err != nil {
				w.handleTxError(tx, req.TxUUID, err)
			}
			w.txs[req.TxUUID] = tx

			bytes, _ := json.Marshal(models.Response{
				Status: successStatus,
				Uuid:   req.TxUUID,
			})
			w.p.SendMessage(&sarama.ProducerMessage{
				Topic: w.topicOut,
				Value: sarama.StringEncoder(bytes),
			})
			fmt.Println("send")
		}
	}
}

func (w *TxWorker) handleTxError(tx *sql.Tx, txUUID string, err error) {
	tx.Rollback()
	var (
		msg      = internalError
		txStatus = uint32(serverFailureStatus)
	)
	if st, ok := status.FromError(err); ok {
		msg = st.Message()
		txStatus = userFailureStatus
	}

	bytes, _ := json.Marshal(txResponse{
		Status:  txStatus,
		Uuid:    txUUID,
		Message: msg,
	})
	w.p.SendMessage(&sarama.ProducerMessage{
		Topic: w.topicOut,
		Value: sarama.StringEncoder(bytes),
	})
}
