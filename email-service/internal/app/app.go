package app

import (
	"context"
	"github.com/alserov/device-shop/email-service/internal/worker"
	"github.com/alserov/device-shop/email-service/internal/worker/order"

	"github.com/alserov/device-shop/email-service/internal/config"
	"github.com/alserov/device-shop/email-service/internal/email"
	"github.com/alserov/device-shop/email-service/internal/logger"
	"github.com/alserov/device-shop/email-service/internal/worker/auth"

	"log/slog"
	"os/signal"
	"syscall"
)

type app struct {
	log *slog.Logger
	cfg *config.Config
}

type App interface {
	MustStart()
}

func NewApp(cfg *config.Config) App {
	return &app{
		log: logger.MustSetupLogger(cfg.Env),
		cfg: cfg,
	}
}

func (a *app) MustStart() {
	defer a.recover()
	ctx := context.Background()
	signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)

	a.mustStartEmailManager(ctx)
}

func (a *app) mustStartEmailManager(ctx context.Context) {
	poster := email.NewEmailManager(a.cfg.Email.Password, a.cfg.Email.Email, a.cfg.Email.Name)

	go auth.StartWorker(&worker.Worker{
		Ctx:        ctx,
		Topic:      a.cfg.Broker.Topics.AuthTopic,
		BrokerAddr: a.cfg.Broker.Addr,
		Poster:     poster,
		Log:        a.log,
	})
	go order.StartWorker(&worker.Worker{
		Ctx:        ctx,
		Topic:      a.cfg.Broker.Topics.OrderTopic,
		BrokerAddr: a.cfg.Broker.Addr,
		Poster:     poster,
		Log:        a.log,
	})

	a.log.Info("app is running")
	<-ctx.Done()
	a.log.Info("app was stopped")
}

func (a *app) recover() {
	err := recover()
	a.log.Error("app was stopped due to panic", slog.Any("error", err))
}
