package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env      string     `yaml:"env"`
	DB       DBConfig   `yaml:"db"`
	GRPC     GRPCConfig `yaml:"grpc"`
	Services Services   `yaml:"services"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type Services struct {
	DeviceAddr string `yaml:"device"`
}

type DBConfig struct {
	Uri string `yaml:"uri"`
}

func MustLoad() *Config {
	path := fetchConfigPath()

	if _, err := os.Stat(path); err != nil {
		panic("config file not found: " + path)
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
		os.Getenv("CONFIG_PATH")
	}

	return path
}
