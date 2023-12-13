package mail

import (
	"fmt"
	"github.com/IBM/sarama"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Emailer interface {
	Send(email string) error
}

func NewEmailer(brokerAddr string, topic string, p sarama.SyncProducer) Emailer {
	return &email{
		p:          p,
		brokerAddr: brokerAddr,
		topic:      topic,
	}
}

type email struct {
	brokerAddr string
	topic      string

	p sarama.SyncProducer
}

const (
	kafkaClientID = "SIGNUP_RPC"
)

func (e *email) Send(email string) error {
	_, _, err := e.p.SendMessage(&sarama.ProducerMessage{
		Value: sarama.StringEncoder(email),
		Topic: e.topic,
	})
	if err != nil {
		return status.Error(codes.Internal, fmt.Sprintf("failed to send message to topic: %v", err))
	}

	return nil
}
