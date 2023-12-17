package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env      string   `yaml:"env"`
	Server   Server   `yaml:"server"`
	Cache    Cache    `yaml:"cache"`
	Services Services `yaml:"services"`
	Broker   Broker   `yaml:"broker"`
}

type Broker struct {
	Addr   string `yaml:"addr"`
	Topics struct {
		Metrics struct {
			UsersAmount string `yaml:"usersAmount"`
			Orders      string `yaml:"orders"`
			Latency     string `yaml:"latency"`
		} `yaml:"metrics"`
	} `yaml:"topics"`
}

type Server struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type Services struct {
	UserAddr   string `yaml:"userAddr"`
	DeviceAddr string `yaml:"deviceAddr"`
	OrderAddr  string `yaml:"orderAddr"`
	CollAddr   string `yaml:"collAddr"`
}

type Cache struct {
	Addr string `yaml:"addr"`
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
