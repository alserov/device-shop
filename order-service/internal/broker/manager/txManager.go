package manager

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/order-service/internal/broker"
	"github.com/alserov/device-shop/order-service/internal/broker/manager/models"

	"github.com/alserov/device-shop/order-service/internal/broker/producer"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"sync"
)

type txManager struct {
	log *slog.Logger

	broker *broker.Broker

	p sarama.SyncProducer
}

type TxManager interface {
	DoTx(in models.TxBody) error
}

const (
	kafkaClientID = "TX_MANAGER"
)

func NewTxManager(b *broker.Broker, log *slog.Logger) TxManager {
	prod, err := producer.NewProducer([]string{b.BrokerAddr}, kafkaClientID)
	if err != nil {
		panic("failed to create producer: " + err.Error())
	}

	return &txManager{
		log:    log,
		p:      prod,
		broker: b,
	}
}

const (
	userFailureStatus = 1
	successStatus     = 2

	txAmount = 2

	internalError = "internal error"
)

func (t *txManager) DoTx(in models.TxBody) error {
	txUUID := uuid.New().String()

	var (
		wg    = &sync.WaitGroup{}
		chErr = make(chan error, 1)
	)
	wg.Add(txAmount)

	go func() {
		defer wg.Done()
		defer fmt.Println("tx 1")
		err := t.startTx(t.broker.Topics.User.In, t.broker.Topics.User.Out, models.BalanceReq{
			TxUUID:     txUUID,
			OrderPrice: in.OrderPrice,
			UserUUID:   in.UserUUID,
		}, txUUID)
		if err != nil {
			t.log.Error("failed to execute balance tx", slog.String("error", err.Error()))
			chErr <- err
		}
	}()

	go func() {
		defer wg.Done()
		defer fmt.Println("tx 2")
		err := t.startTx(t.broker.Topics.Device.In, t.broker.Topics.Device.Out, models.DeviceReq{
			OrderDevices: in.OrderDevices,
			TxUUID:       txUUID,
		}, txUUID)
		if err != nil {
			t.log.Error("failed to execute device tx", slog.String("error", err.Error()))
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

	go func() {
		wg.Wait()
		close(chErr)
	}()

	for err := range chErr {
		t.log.Error("tx error", slog.String("error", err.Error()))
		t.notifyWorkers(userFailureStatus, txUUID)
		return err
	}

	fmt.Println("not")
	t.notifyWorkers(successStatus, txUUID)
	return nil
}

func (t *txManager) startTx(topicIn string, topicOut string, body interface{}, txUUID string) error {
	cons, err := sarama.NewConsumer([]string{t.broker.BrokerAddr}, nil)
	if err != nil {
		return status.Error(codes.Internal, internalError)
	}
	defer cons.Close()

	bytes, err := json.Marshal(body)
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

	msgs, err := t.subscribe(topicOut, cons)
	if err != nil {
		return err
	}

	for msg := range msgs {
		var m models.Response
		if err = json.Unmarshal(msg, &m); err != nil {
			return err
		}

		if txUUID == m.UUID {
			switch m.Status {
			case successStatus:
				return nil
			case userFailureStatus:
				return status.Error(codes.Canceled, m.Message)
			default:
				return status.Error(codes.Internal, internalError)
			}
		}
	}

	return nil
}

func (t *txManager) subscribe(topic string, c sarama.Consumer) (<-chan []byte, error) {
	op := "txManager.subscribe"

	pList, err := c.Partitions(topic)
	if err != nil {
		return nil, err
	}
	offset := sarama.OffsetNewest

	var (
		chMessages = make(chan []byte, 5)
	)
	go func() {
		defer close(chMessages)
		for _, p := range pList {
			pConsumer, err := c.ConsumePartition(topic, p, offset)
			if err != nil {
				t.log.Error("failed to consume partition", slog.String("error", err.Error()), slog.String("op", op))
				return
			}
			for msg := range pConsumer.Messages() {
				chMessages <- msg.Value
			}
		}
	}()

	return chMessages, nil
}

func (t *txManager) notifyWorkers(txStatus uint32, txUUID string) {
	bytes, _ := json.Marshal(models.BalanceReq{
		TxUUID: txUUID,
		Status: txStatus,
	})
	_, _, err := t.p.SendMessage(&sarama.ProducerMessage{
		Topic: t.broker.Topics.User.In,
		Value: sarama.StringEncoder(bytes),
	})
	t.handleSendMessageError(err, t.broker.Topics.User.In)

	bytes, _ = json.Marshal(models.DeviceReq{
		TxUUID: txUUID,
		Status: txStatus,
	})
	_, _, err = t.p.SendMessage(&sarama.ProducerMessage{
		Topic: t.broker.Topics.Device.In,
		Value: sarama.StringEncoder(bytes),
	})
	t.handleSendMessageError(err, t.broker.Topics.Device.In)
}

func (t *txManager) handleSendMessageError(err error, topic string) {
	if err != nil {
		t.log.Error("failed to send message", slog.String("error", err.Error()), slog.String("topic", topic))
	}
}
