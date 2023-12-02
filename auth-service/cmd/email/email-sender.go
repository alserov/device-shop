package main

import (
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/auth-service/internal/broker"
	"github.com/alserov/device-shop/auth-service/internal/utils"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env file not found")
	}

	cons, err := sarama.NewConsumer([]string{os.Getenv("BROKER_ADDR")}, nil)
	if err != nil {
		log.Fatal("failed to create a consumer: ", err)
	}

	msgs, err := broker.Subscribe(os.Getenv("TOPIC"), cons)
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
