package worker

import (
	"context"
	"github.com/alserov/device-shop/email-service/internal/email"
	"log/slog"
)

type Worker struct {
	Ctx        context.Context
	Topic      string
	BrokerAddr string
	Poster     email.Poster
	Log        *slog.Logger
}
