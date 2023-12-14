package workers

import (
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/metrics/internal/broker"
	"github.com/alserov/device-shop/metrics/internal/logger"
	"github.com/alserov/device-shop/metrics/internal/metric/counter"
	"github.com/prometheus/client_golang/prometheus"
	"log/slog"
)

type requestWorker struct {
	log *slog.Logger

	c sarama.Consumer

	count counter.Counter

	topics *broker.Topics
}

func NewRequestWorker(brokerAddr string, topics *broker.Topics, log *slog.Logger) (Worker, []prometheus.Collector) {
	c, err := sarama.NewConsumer([]string{brokerAddr}, sarama.NewConfig())
	if err != nil {
		panic("failed to init consumer: " + err.Error())
	}

	count := counter.NewCounter()

	return &requestWorker{
		log:    log,
		c:      c,
		topics: topics,
		count:  count,
	}, count.Metrics()
}

func (r *requestWorker) Start() {
	go r.startUsersCounter()

	select {}
}

func (r *requestWorker) startUsersCounter() {
	msgs, err := broker.Subscribe(r.topics.User, r.c)
	go func() {
		for e := range err {
			r.log.Error("consumer failed", logger.Error(e, ""))
		}
	}()

	for _ = range msgs.Messages() {
		r.count.IncUsers()
	}
}
