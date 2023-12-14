package workers

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/metrics/internal/broker"
	"github.com/alserov/device-shop/metrics/internal/logger"
	"github.com/alserov/device-shop/metrics/internal/metric/histogram"
	"github.com/prometheus/client_golang/prometheus"
	"log/slog"
	"time"
)

type latencyWorker struct {
	log *slog.Logger

	c sarama.Consumer

	history histogram.History

	topics *broker.Topics
}

func NewLatencyWorker(brokerAddr string, t *broker.Topics, log *slog.Logger) (Worker, []prometheus.Collector) {
	cons, err := sarama.NewConsumer([]string{brokerAddr}, sarama.NewConfig())
	if err != nil {
		panic("failed to init consumer: " + err.Error())
	}

	history := histogram.NewHistory()

	return &latencyWorker{
		log:     log,
		c:       cons,
		history: history,
		topics:  t,
	}, history.Metrics()
}

func (l *latencyWorker) Start() {
	msgs, err := broker.Subscribe(l.topics.Latency, l.c)
	go func() {
		for e := range err {
			l.log.Error("consumer failed", logger.Error(e, ""))
		}
	}()

	for msg := range msgs.Messages() {
		var latency time.Duration
		if err := json.Unmarshal(msg.Value, &latency); err != nil {
			l.log.Error("failed to unmarshall message value", slog.String("error", err.Error()))
			continue
		}

		l.history.UpdateLatency(latency.Seconds())
	}
}
