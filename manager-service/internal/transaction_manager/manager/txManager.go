package manager

import (
	"encoding/json"
	"errors"
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/manager-service/internal/models"
	"github.com/alserov/device-shop/manager-service/internal/transaction_manager/consumer"
	"github.com/alserov/device-shop/manager-service/internal/transaction_manager/producer"
	"github.com/alserov/device-shop/manager-service/internal/utils/converter"
	"github.com/google/uuid"
	"log/slog"
	"sync"
)

type TxManager struct {
	log  *slog.Logger
	conv *converter.Converter

	topics serviceTopics

	p sarama.SyncProducer
	c sarama.Consumer
}

type serviceTopics struct {
	orderService      string
	userService       string
	deviceService     string
	collectionService string
}

const (
	kafkaClientID = "TX_MANAGER"
)

func NewTxManager(brokerAddr string, log *slog.Logger) *TxManager {
	cons, err := sarama.NewConsumer([]string{brokerAddr}, nil)
	if err != nil {
		panic("failed to create consumer: " + err.Error())
	}

	prod, err := producer.NewProducer([]string{brokerAddr}, kafkaClientID)
	if err != nil {
		panic("failed to create producer: " + err.Error())
	}

	return &TxManager{
		p:    prod,
		c:    cons,
		log:  log,
		conv: converter.NewConverter(),
	}
}

const (
	failureStatus = 0
	successStatus = 1
	allTxs        = 4
)

func (t *TxManager) StartController(topic string) {
	msgs, err := consumer.Subscribe(topic, t.c)
	if err != nil {
		panic("failed to subscribe for a topic: " + err.Error())
	}

	for m := range msgs {
		txUUID := uuid.New().String()
		var req models.TxRequest
		if err = json.Unmarshal(m, &req); err != nil {
			t.log.Error("failed to unmarshall message: " + err.Error())
			t.p.SendMessage(&sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.StringEncoder(t.conv.TxResponseToBytes(txUUID, failureStatus, "failed to unmarshall message")),
			})
			continue
		}

		tx := &models.Tx{
			Uuid:   txUUID,
			AllTxs: allTxs,
		}

		var (
			wg    = &sync.WaitGroup{}
			chErr = make(chan error, 1)
		)
		wg.Add(int(tx.AllTxs))

		go func() {
			defer wg.Done()
			if err = t.startTx(t.topics.userService, txUUID); err != nil {
				t.log.Error("failed to execute balance tx", slog.String("error", err.Error()))
				chErr <- err
			}
		}()

		go func() {
			defer wg.Done()
			if err = t.startTx(t.topics.deviceService, txUUID); err != nil {
				t.log.Error("failed to execute device tx", slog.String("error", err.Error()))
				chErr <- err
			}
		}()

		go func() {
			defer wg.Done()
			if err = t.startTx(t.topics.collectionService, txUUID); err != nil {
				t.log.Error("failed to execute collection tx", slog.String("error", err.Error()))
				chErr <- err
			}
		}()

		go func() {
			wg.Wait()
			close(chErr)
		}()

		for _ = range chErr {
			t.handleTxError(txUUID, topic)
			continue
		}

		t.handleSuccess(txUUID, topic)
	}
}

func (t *TxManager) startTx(topic string, txUUID string) error {
	_, _, err := t.p.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(txUUID),
	})
	if err != nil {
		return err
	}

	msgs, err := consumer.Subscribe(topic, t.c)
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
			default:
				return errors.New("tx failed")
			}
		}
	}

	return nil
}

func (t *TxManager) handleSuccess(txUUID string, topic string) {
	t.p.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(t.conv.TxResponseToBytes(txUUID, successStatus, "")),
	})
}

func (t *TxManager) handleTxError(txUUID string, topic string) {
	t.p.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(t.conv.TxResponseToBytes(txUUID, failureStatus, "")),
	})
}
