package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env      string       `yaml:"env" env-default:"local"`
	DB       DBConfig     `yaml:"db"`
	GRPC     GRPCConfig   `yaml:"grpc"`
	Broker   BrokerConfig `yaml:"broker"`
	Services Services     `yaml:"services"`
}

type Services struct {
	CollectionAddr string `yaml:"collection"`
}

type BrokerConfig struct {
	Addr           string `yaml:"addr"`
	WorkerTopicIn  string `yaml:"workerTopicIn"`
	WorkerTopicOut string `yaml:"workerTopicOut"`
}

type DBConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	SSLMode  string `yaml:"sslmode"`
	Host     string `yaml:"host"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); err != nil {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var path string

	flag.StringVar(&path, "c", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
