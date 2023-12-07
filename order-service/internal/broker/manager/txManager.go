package manager

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/order-service/internal/broker/producer"
	"github.com/alserov/device-shop/order-service/internal/service/models"
	"github.com/alserov/device-shop/order-service/internal/utils/converter"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"sync"
)

type txManager struct {
	log  *slog.Logger
	conv *converter.BrokerConverter

	topics     serviceTopics
	brokerAddr string

	p sarama.SyncProducer
}

type serviceTopics struct {
	userInTopic     string
	userOutTopic    string
	deviceTopic     string
	collectionTopic string
}

type TxManager interface {
	DoTx(in models.DoTxBody) error
}

const (
	kafkaClientID = "TX_MANAGER"
)

func NewTxManager(brokerAddr string, deviceTopic string, userInTopic string, userOutTopic string, collectionTopic string, log *slog.Logger) TxManager {
	prod, err := producer.NewProducer([]string{brokerAddr}, kafkaClientID)
	if err != nil {
		panic("failed to create producer: " + err.Error())
	}

	return &txManager{
		log:        log,
		p:          prod,
		brokerAddr: brokerAddr,
		conv:       converter.NewBrokerConverter(),
		topics: serviceTopics{
			deviceTopic:     deviceTopic,
			userInTopic:     userInTopic,
			userOutTopic:    userOutTopic,
			collectionTopic: collectionTopic,
		},
	}
}

const (
	userFailureStatus = 0
	successStatus     = 1
	txAmount          = 1

	internalError = "internal error"
)

func (t *txManager) DoTx(in models.DoTxBody) error {
	tx := &models.Tx{
		Uuid: uuid.New().String(),
	}

	var (
		wg    = &sync.WaitGroup{}
		chErr = make(chan error, 1)
	)
	wg.Add(txAmount)

	go func() {
		defer wg.Done()
		if err := t.startTx(t.topics.userInTopic, t.topics.userOutTopic, models.TxBalanceReq{
			TxUUID:     tx.Uuid,
			OrderPrice: in.OrderPrice,
			UserUUID:   in.UserUUID,
		}, tx.Uuid); err != nil {
			t.log.Error("failed to execute balance tx", slog.String("error", err.Error()))
			chErr <- err
		}
	}()

	//go func() {
	//	defer wg.Done()
	//	if err := t.startTx(t.topics.deviceService, models.TxDeviceReq{OrderDevices: in.OrderDevices}, tx.Uuid); err != nil {
	//		t.log.Error("failed to execute device tx", slog.String("error", err.Error()))
	//		chErr <- err
	//	}
	//}()

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
		t.notifyTxs(models.TxBalanceReq{
			TxUUID: tx.Uuid,
			Status: userFailureStatus,
		})
		return err
	}

	t.notifyTxs(models.TxBalanceReq{
		TxUUID: tx.Uuid,
		Status: successStatus,
	})
	return nil
}

func (t *txManager) startTx(topicIn string, topicOut string, body interface{}, txUUID string) error {
	cons, err := sarama.NewConsumer([]string{t.brokerAddr}, nil)
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
		var m models.TxResponse
		if err = json.Unmarshal(msg, &m); err != nil {
			return err
		}
		if txUUID == m.Uuid {
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

func (t *txManager) notifyTxs(body interface{}) {
	bytes, _ := json.Marshal(body)

	_, _, err := t.p.SendMessage(&sarama.ProducerMessage{
		Topic: t.topics.userInTopic,
		Value: sarama.StringEncoder(bytes),
	})
	t.handleSendMessageError(err, t.topics.userInTopic)

	//	_, _, err = t.p.SendMessage(&sarama.ProducerMessage{
	//		Topic: t.topics.collectionService,
	//		Value: sarama.StringEncoder(bytes),
	//	})
	//	t.handleSendMessageError(err, t.topics.collectionService)
	//
	//	_, _, err = t.p.SendMessage(&sarama.ProducerMessage{
	//		Topic: t.topics.deviceService,
	//		Value: sarama.StringEncoder(bytes),
	//	})
	//	t.handleSendMessageError(err, t.topics.deviceService)
}

func (t *txManager) handleSendMessageError(err error, topic string) {
	if err != nil {
		t.log.Error("failed to send message", slog.String("error", err.Error()), slog.String("topic", topic))
	}
}
