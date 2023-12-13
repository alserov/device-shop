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

	topics *broker.RequestTopics
}

func NewRequestWorker(brokerAddr string, topics *broker.RequestTopics) (Worker, prometheus.Collector) {
	c, err := sarama.NewConsumer([]string{brokerAddr}, sarama.NewConfig())
	if err != nil {
		panic("failed to init consumer: " + err.Error())
	}

	count := request.NewCounter()

	return &requestWorker{
		c:      c,
		topics: topics,
		count:  count,
	}, count.Metric()
}

func (r *requestWorker) Start() {
	go r.startTotalCounter()
	go r.startSuccessfulCounter()

	select {}
}

func (r *requestWorker) startTotalCounter() {
	msgs, err := broker.Subscribe(r.topics.Total, r.c)
	go func() {
		for e := range err {
			r.log.Error("consumer failed", logger.Error(e, ""))
		}
	}()

	for _ = range msgs.Messages() {
		r.count.Inc()
	}
}

func (r *requestWorker) startSuccessfulCounter() {
	msgs, err := broker.Subscribe(r.topics.Successful, r.c)
	go func() {
		for e := range err {
			r.log.Error("consumer failed", logger.Error(e, ""))
		}
	}()

	for _ = range msgs.Messages() {
		r.count.Inc()
	}
}
