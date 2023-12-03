package main

import (
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/email-service/internal/config"
	"github.com/alserov/device-shop/email-service/internal/consumer"
	"github.com/alserov/device-shop/email-service/internal/email"
	"github.com/alserov/device-shop/email-service/internal/logger"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoadEmail()

	log := logger.MustSetupLogger(cfg.Env)

	poster := email.NewPoster(cfg.Email.Password, cfg.Email.Email, cfg.Email.Name)

	cons, err := sarama.NewConsumer([]string{cfg.Kafka.BrokerAddr}, nil)
	if err != nil {
		panic("failed to create a consumer: " + err.Error())
	}

	log.Info("email service started", slog.String("env", cfg.Env))
	go func() {
		msgs, err := consumer.Subscribe(cfg.Kafka.Topics.Auth, cons)
		if err != nil {
			panic("failed to subscribe for a topic: " + err.Error())
		}
		for i := 0; i < 5; i++ {
			go func() {
				for m := range msgs {
					if err = poster.SendAuth(m); err != nil {
						log.Error("failed to send message", slog.String("error", err.Error()))
					}
				}
			}()
		}
	}()

	go func() {
		msgs, err := consumer.Subscribe(cfg.Kafka.Topics.Order, cons)
		if err != nil {
			panic("failed to subscribe for a topic: " + err.Error())
		}
		for i := 0; i < 5; i++ {
			go func() {
				for m := range msgs {
					if err = poster.SendOrder(m); err != nil {
						log.Error("failed to send message", slog.String("error", err.Error()))
					}
				}
			}()
		}
	}()

	chStop := make(chan os.Signal, 1)
	signal.Notify(chStop, syscall.SIGINT, syscall.SIGTERM)
	sign := <-chStop
	log.Info("server has stopped", slog.String("signal", sign.String()))
}
