package main

import (
	"context"
	"github.com/alserov/shop/user-service/internal/app"
	"os"
	"os/signal"
)

func main() {
	a, err := app.New()
	if err != nil {
		panic(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err = a.Start(ctx); err != nil {
		panic(err)
	}
}
