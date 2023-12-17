package order

import (
	"github.com/IBM/sarama"

	"github.com/alserov/device-shop/email-service/internal/broker"
	"github.com/alserov/device-shop/email-service/internal/worker"

	"log/slog"
)

func StartWorker(w *worker.Worker) {
	cons, err := sarama.NewConsumer([]string{w.BrokerAddr}, nil)
	if err != nil {
		panic("failed to create a consumer: " + err.Error())
	}

	msgs, err := broker.Subscribe(w.Topic, cons)
	if err != nil {
		panic("failed to subscribe for a topic: " + err.Error())
	}
	for i := 0; i < 5; i++ {
		go func() {
			for m := range msgs {
				if err = w.Poster.SendOrder(string(m)); err != nil {
					w.Log.Error("failed to send message", slog.String("error", err.Error()))
				}
			}
		}()
	}
	select {
	case <-w.Ctx.Done():
	}
}
