package manager

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/order-service/internal/service/models"

	"github.com/alserov/device-shop/order-service/internal/broker"
	brokermodels "github.com/alserov/device-shop/order-service/internal/broker/manager/models"
	repo "github.com/alserov/device-shop/order-service/internal/db/models"
	"github.com/alserov/device-shop/order-service/internal/utils/converter"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"sync"
	"time"
)

type txManager struct {
	log *slog.Logger

	broker *broker.Broker

	conv converter.ServiceConverter

	p sarama.SyncProducer
}

type TxManager interface {
	CreateOrderTx(in brokermodels.CreateOrderTxBody) error
	CancelOrderTx(in brokermodels.CancelOrderTxBody) error
}

const (
	kafkaClientID = "TX_MANAGER"
)

func NewTxManager(b *broker.Broker, log *slog.Logger) TxManager {
	prod, err := broker.NewProducer([]string{b.Addr}, kafkaClientID)
	if err != nil {
		panic("failed to create producer: " + err.Error())
	}

	return &txManager{
		log:    log,
		p:      prod,
		broker: b,
		conv:   converter.NewServiceConverter(),
	}
}

const (
	serverFailureStatus = 0
	userFailureStatus   = 1
	successStatus       = 2

	// createOrderTxAmount - number of services to create order
	createOrderTxAmount = 3
	// cancelOrderTxAmount - number of services to cancel order
	cancelOrderTxAmount = 2

	internalError = "internal error"
)

func (t *txManager) CancelOrderTx(in brokermodels.CancelOrderTxBody) error {
	txUUID := uuid.New().String()

	var (
		wg    = &sync.WaitGroup{}
		chErr = make(chan error)
	)

	wg.Add(cancelOrderTxAmount)

	go func() {
		defer wg.Done()
		err := t.startTx(t.broker.Topics.DeviceRollback.In, t.broker.Topics.Device.Out, brokermodels.DeviceReq[repo.OrderDevice]{
			TxUUID:       txUUID,
			OrderDevices: in.OrderDevices,
		}, txUUID)
		if err != nil {
			chErr <- err
		}
	}()

	go func() {
		defer wg.Done()
		err := t.startTx(t.broker.Topics.BalanceRefund.In, t.broker.Topics.BalanceRefund.Out, brokermodels.BalanceReq{
			TxUUID:     txUUID,
			OrderPrice: in.OrderPrice,
			UserUUID:   in.UserUUID,
		}, txUUID)
		if err != nil {
			chErr <- err
		}
	}()

	go func() {
		wg.Wait()
		close(chErr)
	}()

	for err := range chErr {
		t.notifyWorkers(userFailureStatus, txUUID)
		return err
	}

	t.notifyWorkers(successStatus, txUUID)
	return nil
}

func (t *txManager) CreateOrderTx(in brokermodels.CreateOrderTxBody) error {
	txUUID := uuid.New().String()

	var (
		wg    = &sync.WaitGroup{}
		chErr = make(chan error, createOrderTxAmount)
		tx    *sql.Tx
	)
	wg.Add(createOrderTxAmount)

	go func() {
		defer wg.Done()
		err := t.startTx(t.broker.Topics.Balance.In, t.broker.Topics.Balance.Out, brokermodels.BalanceReq{
			TxUUID:     txUUID,
			OrderPrice: in.OrderPrice,
			UserUUID:   in.UserUUID,
		}, txUUID)
		if err != nil {
			chErr <- err
		}
	}()

	go func() {
		defer wg.Done()
		err := t.startTx(t.broker.Topics.Device.In, t.broker.Topics.Device.Out, brokermodels.DeviceReq[models.OrderDevice]{
			OrderDevices: in.OrderDevices,
			TxUUID:       txUUID,
		}, txUUID)
		if err != nil {
			chErr <- err
		}
	}()

	go func() {
		defer wg.Done()

		var err error
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
		defer cancel()

		tx, err = in.Repo.CreateOrderTx(ctx, t.conv.CreateOrderReqToRepo(in.Order, in.OrderUUID))
		if err != nil {
			chErr <- err
		}
	}()

	//go func() {
	//	defer wg.Done()
	//	if err := t.startTx(t.topics.collectionService, tx.Uuid); err != nil {
	//		t.log.Error("failed to execute collection tx", slog.String("error", err.Error()))
	//		chErr <- err
	//	}
	//}()

	wg.Wait()
	close(chErr)

	for err := range chErr {
		t.notifyWorkers(userFailureStatus, txUUID)
		if err := tx.Rollback(); err != nil {
			t.log.Error("failed to rollback", slog.String("error", err.Error()), slog.String("op", "txManager.CreateOrderTx"))
			return status.Error(codes.Internal, internalError)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		t.notifyWorkers(serverFailureStatus, txUUID)
		return status.Error(codes.Internal, internalError)
	}
	t.notifyWorkers(successStatus, txUUID)
	return nil
}

func (t *txManager) startTx(topicIn string, topicOut string, body interface{}, txUUID string) error {
	cons, err := sarama.NewConsumer([]string{t.broker.Addr}, nil)
	if err != nil {
		return status.Error(codes.Internal, internalError)
	}
	defer cons.Close()

	bytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	pConsumer, err := t.subscribe(topicOut, cons)
	if err != nil {
		return err
	}

	_, _, err = t.p.SendMessage(&sarama.ProducerMessage{
		Topic: topicIn,
		Value: sarama.StringEncoder(bytes),
	})
	if err != nil {
		return err
	}

	for msg := range pConsumer.Messages() {
		var res brokermodels.Response
		if err = json.Unmarshal(msg.Value, &res); err != nil {
			return err
		}

		if txUUID == res.UUID {
			switch res.Status {
			case successStatus:
				return nil
			case userFailureStatus:
				return status.Error(codes.Canceled, res.Message)
			default:
				return status.Error(codes.Internal, internalError)
			}
		}
	}

	return nil
}

func (t *txManager) subscribe(topic string, c sarama.Consumer) (sarama.PartitionConsumer, error) {
	op := "txManager.subscribe"

	pList, err := c.Partitions(topic)
	if err != nil {
		return nil, err
	}
	offset := sarama.OffsetNewest

	var pConsumer sarama.PartitionConsumer
	for _, p := range pList {
		pConsumer, err = c.ConsumePartition(topic, p, offset)
		if err != nil {
			t.log.Error("failed to consume partition", slog.String("error", err.Error()), slog.String("op", op))
			return nil, status.Error(codes.Internal, internalError)
		}
	}

	return pConsumer, nil
}

func (t *txManager) notifyWorkers(txStatus uint32, txUUID string) {
	bytes, _ := json.Marshal(brokermodels.BalanceReq{
		TxUUID: txUUID,
		Status: txStatus,
	})
	_, _, err := t.p.SendMessage(&sarama.ProducerMessage{
		Topic: t.broker.Topics.Balance.In,
		Value: sarama.StringEncoder(bytes),
	})
	t.handleSendMessageError(err, t.broker.Topics.Balance.In)

	bytes, _ = json.Marshal(brokermodels.DeviceReq[any]{
		TxUUID: txUUID,
		Status: txStatus,
	})
	_, _, err = t.p.SendMessage(&sarama.ProducerMessage{
		Topic: t.broker.Topics.Device.In,
		Value: sarama.StringEncoder(bytes),
	})
	t.handleSendMessageError(err, t.broker.Topics.Device.In)

	if txStatus == serverFailureStatus {
		t.log.Error("failed to execute tx", slog.String("error", err.Error()))
	}
}

func (t *txManager) handleSendMessageError(err error, topic string) {
	if err != nil {
		t.log.Error("failed to send message", slog.String("error", err.Error()), slog.String("topic", topic))
	}
}
