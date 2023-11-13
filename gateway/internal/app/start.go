package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (a *App) Start(ctx context.Context) error {
	log.Println("Starting API Gateway")
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", a.port),
		Handler:      a.router,
		WriteTimeout: time.Duration(a.timeout) * time.Second,
		ReadTimeout:  time.Duration(a.timeout) * time.Second,
	}

	chErr := make(chan error, 1)
	go func() {
		log.Println("API Gateway is working")
		if err := srv.ListenAndServe(); err != nil {
			chErr <- err
		}
	}()

	select {
	case <-ctx.Done():
		err := srv.Shutdown(ctx)
		if err != nil {
			return err
		}
		return nil
	case err := <-chErr:
		return err
	}
}

type Config struct {
	Server struct {
		Port    int    `yaml:"port"`
		Host    string `yaml:"host"`
		Timeout int    `yaml:"timeout"`
	} `yaml:"server"`
}
