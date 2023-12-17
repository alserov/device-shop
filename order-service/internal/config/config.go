package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env      string       `yaml:"env" env-default:"local"`
	DB       DBConfig     `yaml:"db"`
	GRPC     GRPCConfig   `yaml:"grpc"`
	Broker   BrokerConfig `yaml:"broker"`
	Services Services     `json:"services"`
}

type Services struct {
	DeviceAddr string `yaml:"deviceAddr"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password int    `yaml:"password"`
	Name     string `yaml:"name"`
	Sslmode  string `yaml:"sslmode"`
}

type GRPCConfig struct {
	Port    int    `yaml:"port"`
	Timeout string `yaml:"timeout"`
}

type BrokerConfig struct {
	Addr string `yaml:"brokerAddr"`

	UserTopicIn  string `yaml:"userTopicIn"`
	UserTopicOut string `yaml:"userTopicOut"`

	DeviceTopicIn  string `yaml:"deviceTopicIn"`
	DeviceTopicOut string `yaml:"deviceTopicOut"`

	CollectionTopicIn  string `yaml:"collectionTopicIn"`
	CollectionTopicOut string `yaml:"collectionTopicOut"`
}

func MustLoad() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("config path is empty")
	}

	_, err := os.Stat(path)
	if err != nil {
		panic("config file not found: " + path)
	}

	var cfg Config
	if err = cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var path string

	flag.StringVar(&path, "c", "", "path to config file")
	flag.Parse()

	if path == "" {
		os.Getenv("CONFIG_PATH")
	}

	return path
}
