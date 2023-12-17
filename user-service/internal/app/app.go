package app

import (
	"fmt"
	"github.com/alserov/device-shop/user-service/internal/broker"
	"github.com/alserov/device-shop/user-service/internal/broker/worker"
	"github.com/alserov/device-shop/user-service/internal/config"
	"github.com/alserov/device-shop/user-service/internal/db/postgres"
	"github.com/alserov/device-shop/user-service/internal/logger"
	"github.com/alserov/device-shop/user-service/internal/server"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type app struct {
	port  int
	dbDsn string

	log        *slog.Logger
	gRPCServer *grpc.Server
	broker     *broker.Broker
}

type App interface {
	MustStart()
}

func New(cfg *config.Config) App {
	return &app{
		port:       cfg.GRPC.Port,
		log:        logger.MustSetupLogger(cfg.Env),
		gRPCServer: grpc.NewServer(),
		broker: &broker.Broker{
			Addr: cfg.Kafka.BrokerAddr,
			Topics: &broker.Topics{
				Email: cfg.Kafka.EmailTopic,
				Worker: broker.WorkerTopic{
					In:  cfg.Kafka.WorkerTopicIn,
					Out: cfg.Kafka.WorkerTopicOut,
				},
			},
		},
		dbDsn: fmt.Sprintf("host=%s port=%d user=%s password=%v dbname=%s sslmode=%s",
			cfg.DB.Host,
			cfg.DB.Port,
			cfg.DB.User,
			cfg.DB.Password,
			cfg.DB.Name,
			cfg.DB.SSLMode,
		),
	}
}

func (a *app) MustStart() {
	defer a.recover()
	a.log.Info("starting app", slog.Int("port", a.port))

	db := postgres.MustConnect(a.dbDsn)
	a.log.Info("db connected")
	repo := postgres.NewRepo(db, a.log)

	w := worker.NewWorker(a.broker, repo, a.log)
	go w.MustStart()

	server.MustRegister(&server.Server{
		GRPCServer: a.gRPCServer,
		Repo:       repo,
		Log:        a.log,
		Broker: &broker.Broker{
			Addr: a.broker.Addr,
			Topics: &broker.Topics{
				Email:  a.broker.Topics.Email,
				Worker: a.broker.Topics.Worker,
			},
		},
	})

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		panic("failed to create listener: " + err.Error())
	}

	chStop := make(chan os.Signal)
	signal.Notify(chStop, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		a.log.Info("app is running")
		if err = a.gRPCServer.Serve(l); err != nil {
			panic("app has stopped due to the error: " + err.Error())
		}
	}()

	sign := <-chStop
	a.stop(sign)
}

func (a *app) stop(sign os.Signal) {
	a.gRPCServer.GracefulStop()
	a.log.Info("app was stopped due to: ", sign.String())
}

func (a *app) recover() {
	err := recover()
	if err != nil {
		a.log.Error("app was stopped due to panic", slog.Any("error", err))
	}
}
