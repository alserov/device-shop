package broker

import (
	"fmt"
	"github.com/IBM/sarama"
)

type requestProducer struct {
	p sarama.SyncProducer

	topics *RequestTopics
}

type RequestProducer interface {
	Inc() error
	IncSuccess() error
}

func NewRequestProducer(p sarama.SyncProducer, t *RequestTopics) RequestProducer {
	return &requestProducer{
		p:      p,
		topics: t,
	}
}

func (rp *requestProducer) Inc() error {
	_, _, err := rp.p.SendMessage(&sarama.ProducerMessage{
		Topic: rp.topics.Total,
	})
	if err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}

	return nil
}

func (rp *requestProducer) IncSuccess() error {
	_, _, err := rp.p.SendMessage(&sarama.ProducerMessage{
		Topic: rp.topics.Successful,
	})
	if err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}

	return nil
}
