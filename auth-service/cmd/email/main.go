package main

import (
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/auth-service/internal/broker"
	"github.com/alserov/device-shop/auth-service/internal/config"
	"github.com/alserov/device-shop/auth-service/internal/utils"
	"log"
)

func main() {
	cfg := config.MustLoadEmail()

	cons, err := sarama.NewConsumer([]string{cfg.Kafka.BrokerAddr}, nil)
	if err != nil {
		log.Fatal("failed to create a consumer: ", err)
	}

	msgs, err := broker.Subscribe(cfg.Kafka.Topic, cons)
	if err != nil {
		log.Fatal("failed to subscribe for a topic: ", err)
	}

	log.Println("email service started")
	for m := range msgs {
		m := m
		go func() {
			if err = utils.SendEmail(m); err != nil {
				log.Println("failed to send message: ", err)
			}
		}()
	}
}
