package broker

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/user-service/internal/broker/producer"
	"github.com/alserov/device-shop/user-service/internal/db/postgres"
	"google.golang.org/grpc/status"
	"log"

	"github.com/alserov/device-shop/user-service/internal/broker/consumer"
	"github.com/alserov/device-shop/user-service/internal/db"
	"github.com/alserov/device-shop/user-service/internal/service/models"
	"github.com/alserov/device-shop/user-service/internal/utils/converter"

	"log/slog"
)

type TxWorker struct {
	log *slog.Logger

	txs  map[string]*sql.Tx
	conv *converter.ServiceConverter

	topicIn  string
	topicOut string

	repo db.UserRepo

	c sarama.Consumer
	p sarama.SyncProducer
}

const (
	internalError = "internal error"
	kafkaClientID = "USER_WORKER"
)

func NewTxWorker(brokerAddr string, topicIn string, topicOut string, db *sql.DB, log *slog.Logger) *TxWorker {
	cons, err := sarama.NewConsumer([]string{brokerAddr}, nil)
	if err != nil {
		panic("failed to start kafka consumer: " + err.Error())
	}

	prod, err := producer.NewProducer([]string{brokerAddr}, kafkaClientID)
	if err != nil {
		panic("failed to start kafka producer: " + err.Error())
	}

	return &TxWorker{
		log:      log,
		c:        cons,
		p:        prod,
		conv:     converter.NewServiceConverter(),
		repo:     postgres.NewRepo(db, log),
		topicIn:  topicIn,
		topicOut: topicOut,
		txs:      make(map[string]*sql.Tx),
	}
}

const (
	serverFailureStatus = 0
	userFailureStatus   = 0
	successStatus       = 1
)

type txResponse struct {
	// 0 - failed
	// 1 - success
	Status  uint32 `json:"status"`
	Message string `json:"message"`
	Uuid    string `json:"uuid"`
}

func (t *TxWorker) MustStart() {
	msgs, err := consumer.Subscribe(t.topicIn, t.c)
	if err != nil {
		panic("failed to subscribe on topic: " + err.Error())
	}

	t.log.Info("worker is running")
	for msg := range msgs {
		var txBalanceReq models.WorkerBalanceReq
		if err = json.Unmarshal(msg, &txBalanceReq); err != nil {
			t.log.Error("failed to unmarshall balance req: " + err.Error())
			continue
		}

		if _, ok := t.txs[txBalanceReq.TxUUID]; ok {
			switch txBalanceReq.Status {
			case successStatus:
				if err := t.txs[txBalanceReq.TxUUID].Commit(); err != nil {
					log.Println(err)
				}
			default:
				t.txs[txBalanceReq.TxUUID].Rollback()
			}
			delete(t.txs, txBalanceReq.TxUUID)
		} else {
			tx, err := t.repo.DebitBalanceTx(context.Background(), t.conv.Balance.WorkerBalanceReqToRepo(txBalanceReq))
			if err != nil {
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
					Uuid:    txBalanceReq.TxUUID,
					Message: msg,
				})
				t.p.SendMessage(&sarama.ProducerMessage{
					Topic: t.topicOut,
					Value: sarama.StringEncoder(bytes),
				})
				continue
			}

			t.txs[txBalanceReq.TxUUID] = tx

			bytes, err := json.Marshal(models.TxResponse{
				Status: successStatus,
				Uuid:   txBalanceReq.TxUUID,
			})
			t.p.SendMessage(&sarama.ProducerMessage{
				Topic: t.topicOut,
				Value: sarama.StringEncoder(bytes),
			})
		}
	}
}
