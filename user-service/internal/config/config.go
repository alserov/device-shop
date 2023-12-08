package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env   string      `yaml:"env"  env-default:"local"`
	GRPC  GRPCConfig  `yaml:"grpc"`
	DB    DBConfig    `yaml:"db"`
	Kafka KafkaConfig `yaml:"kafka"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type KafkaConfig struct {
	BrokerAddr     string `yaml:"brokerAddr"`
	EmailTopic     string `yaml:"emailTopic"`
	WorkerTopicIn  string `yaml:"workerTopicIn"`
	WorkerTopicOut string `yaml:"workerTopicOut"`
}
type DBConfig struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslmode"`
	User     string `yaml:"user"`
}

func MustLoad() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("config path can not be empty")
	}

	if _, err := os.Stat(path); err != nil {
		panic("failed to find config path: " + path)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config file: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var path string

	flag.StringVar(&path, "c", "", "path co config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
