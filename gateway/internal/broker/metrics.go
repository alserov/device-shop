package broker

import (
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

type requestProducer struct {
	p sarama.SyncProducer

	metrics *Metrics
}

type MetricsProducer interface {
	Latency(latency time.Duration) error
	IncUsers() error
	NewOrder() error
}

func NewMetricsProducer(p sarama.SyncProducer, m *Metrics) MetricsProducer {
	return &requestProducer{
		p:       p,
		metrics: m,
	}
}

func (rp *requestProducer) Latency(latency time.Duration) error {
	_, _, err := rp.p.SendMessage(&sarama.ProducerMessage{
		Topic: rp.metrics.Latency,
		Value: sarama.StringEncoder(fmt.Sprintf("%d", latency.Milliseconds())),
	})
	if err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}

	return nil
}

func (rp *requestProducer) IncUsers() error {
	_, _, err := rp.p.SendMessage(&sarama.ProducerMessage{
		Topic: rp.metrics.Users,
	})

	if err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}

	return nil
}

func (rp *requestProducer) NewOrder() error {
	_, _, err := rp.p.SendMessage(&sarama.ProducerMessage{
		Topic: rp.metrics.Orders,
	})

	if err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}

	return nil
}
