package app

import (
	"context"
	"github.com/alserov/device-shop/manager-service/internal/transaction_manager/manager"

	"github.com/alserov/device-shop/manager-service/internal/config"
	"log/slog"
	"os/signal"
	"syscall"
)

type TxApp struct {
	log *slog.Logger
	cfg *config.TxConfig
}

func NewTxApp(cfg *config.TxConfig, log *slog.Logger) *TxApp {
	return &TxApp{
		log: log,
		cfg: cfg,
	}
}

func (a *TxApp) MustStart() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	a.mustStartTxManager(ctx, a.cfg)
}

func (a *TxApp) mustStartTxManager(ctx context.Context, cfg *config.TxConfig) {
	txManager := manager.NewTxManager(a.cfg.BrokerAddr, a.log)
	go txManager.StartController(a.cfg.Topics.Tx.Topic)

	a.log.Info("app is running")
	select {
	case <-ctx.Done():
		a.log.Info("app was stopped")
	}
}
