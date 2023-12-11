package workers

import (
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/metrics/internal/broker"
	"github.com/alserov/device-shop/metrics/internal/logger"
	"github.com/alserov/device-shop/metrics/internal/metric/request"
	"github.com/prometheus/client_golang/prometheus"
	"log/slog"
)

type requestWorker struct {
	log *slog.Logger

	c sarama.Consumer

	count request.Counter

	topic string
}

func NewRequestWorker(brokerAddr string, topic string) (Worker, prometheus.Collector) {
	c, err := sarama.NewConsumer([]string{brokerAddr}, sarama.NewConfig())
	if err != nil {
		panic("failed to init consumer: " + err.Error())
	}

	count := request.NewCounter()

	return &requestWorker{
		c:     c,
		topic: topic,
		count: count,
	}, count.Metric()
}

func (r *requestWorker) Start() {
	msgs, err := broker.Subscribe(r.topic, r.c)
	go func() {
		for e := range err {
			r.log.Error("consumer failed", logger.Error(e, ""))
		}
	}()

	for _ = range msgs.Messages() {
		r.count.Inc()
	}
}
