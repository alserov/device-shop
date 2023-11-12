package main

import (
	"context"
	"github.com/alserov/shop/api/internal/app"
	"os"
	"os/signal"
)

// @title Device Shop
// @version 1.0
// @description API Gateway for Device Shop

// @host localhost:8001
// @BasePath /

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
