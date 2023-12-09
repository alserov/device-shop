package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/user-service/internal/broker"
	"github.com/alserov/device-shop/user-service/internal/broker/consumer"
	"github.com/alserov/device-shop/user-service/internal/broker/producer"
	"github.com/alserov/device-shop/user-service/internal/broker/worker/models"
	"github.com/alserov/device-shop/user-service/internal/db"
	"github.com/alserov/device-shop/user-service/internal/db/postgres"
	"github.com/alserov/device-shop/user-service/internal/utils/converter"
	"google.golang.org/grpc/status"

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

func (w *TxWorker) MustStart() {
	msgs, err := consumer.Subscribe(w.topicIn, w.c)
	if err != nil {
		panic("failed to subscribe on topic: " + err.Error())
	}

	w.log.Info("worker is running")
	for msg := range msgs.Messages() {
		var req models.Request
		if err = json.Unmarshal(msg.Value, &req); err != nil {
			w.log.Error("failed to unmarshall balance req: " + err.Error())
			continue
		}

		if _, ok := w.txs[req.TxUUID]; ok {
			switch req.Status {
			case successStatus:
				if err = w.txs[req.TxUUID].Commit(); err != nil {
					w.log.Error("failed to commit tx", slog.String("error", err.Error()))
				}
			default:
				if err = w.txs[req.TxUUID].Rollback(); err != nil {
					w.log.Error("failed to rollback tx", slog.String("error", err.Error()))
				}
			}
			delete(w.txs, req.TxUUID)
			continue
		}

		tx, err := w.repo.DebitBalanceTx(context.Background(), w.conv.WorkerBalanceReqToRepo(req))
		w.txs[req.TxUUID] = tx
		if err != nil {
			w.handleTxError(req.TxUUID, err)
			continue
		}

		w.sendMessage(req.TxUUID)
	}
}

func (w *TxWorker) sendMessage(txUUID string) {
	bytes, _ := json.Marshal(models.Response{
		Status: successStatus,
		Uuid:   txUUID,
	})
	_, _, err := w.p.SendMessage(&sarama.ProducerMessage{
		Topic: w.topicOut,
		Value: sarama.StringEncoder(bytes),
	})
	if err != nil {
		w.log.Error("failed to send message", slog.String("error", err.Error()))
	}
}

func (w *TxWorker) handleTxError(txUUID string, err error) {
	var (
		msg      = internalError
		txStatus = uint32(serverFailureStatus)
	)
	if st, ok := status.FromError(err); ok {
		msg = st.Message()
		txStatus = userFailureStatus
	}

	bytes, _ := json.Marshal(models.Response{
		Status:  txStatus,
		Uuid:    txUUID,
		Message: msg,
	})

	_, _, err = w.p.SendMessage(&sarama.ProducerMessage{
		Topic: w.topicOut,
		Value: sarama.StringEncoder(bytes),
	})
	if err != nil {
		w.log.Error("failed to send message", slog.String("error", err.Error()))
	}
}
