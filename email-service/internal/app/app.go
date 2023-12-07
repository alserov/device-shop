package app

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/email-service/internal/broker/consumer"
	"github.com/alserov/device-shop/email-service/internal/config"
	"github.com/alserov/device-shop/email-service/internal/email"
	"log/slog"
	"os/signal"
	"syscall"
)

type EmailApp struct {
	log *slog.Logger
	cfg *config.EmailConfig
}

func NewEmailApp(cfg *config.EmailConfig, log *slog.Logger) *EmailApp {
	return &EmailApp{
		log: log,
		cfg: cfg,
	}
}

func (a *EmailApp) MustStartEmail() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go a.mustStartEmailManager(ctx, a.cfg)

	a.log.Info("app is running")
	select {
	case <-ctx.Done():
		a.log.Info("app was stopped")
	}
}

func (a *EmailApp) mustStartEmailManager(ctx context.Context, cfg *config.EmailConfig) {
	poster := email.NewEmailManager(cfg.Email.Password, cfg.Email.Email, cfg.Email.Name)

	cons, err := sarama.NewConsumer([]string{a.cfg.BrokerAddr}, nil)
	if err != nil {
		panic("failed to create a consumer: " + err.Error())
	}

	go func() {
		msgs, err := consumer.Subscribe(cfg.Topics.Email.AuthTopic, cons)
		if err != nil {
			panic("failed to subscribe for a topic: " + err.Error())
		}
		for i := 0; i < 5; i++ {
			go func() {
				for m := range msgs {
					if err = poster.SendAuth(string(m)); err != nil {
						a.log.Error("failed to send message", slog.String("error", err.Error()))
					}
				}
			}()
		}
		select {
		case <-ctx.Done():
		}
	}()

	go func() {
		msgs, err := consumer.Subscribe(cfg.Topics.Email.OrderTopic, cons)
		if err != nil {
			panic("failed to subscribe for a topic: " + err.Error())
		}
		for i := 0; i < 5; i++ {
			go func() {
				for m := range msgs {
					if err = poster.SendOrder(string(m)); err != nil {
						a.log.Error("failed to send message", slog.String("error", err.Error()))
					}
				}
			}()
		}
		select {
		case <-ctx.Done():
		}
	}()

	select {
	case <-ctx.Done():
	}
}
