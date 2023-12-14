package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env    string       `yaml:"env"`
	Server ServerConfig `yaml:"server"`
	Broker BrokerConfig `yaml:"broker"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type BrokerConfig struct {
	Addr   string `yaml:"addr"`
	Topics Topics `yaml:"topics"`
}

type Topics struct {
	UsersAmount string `yaml:"usersAmount"`
	Orders      string `yaml:"orders"`
	Latency     string `yaml:"latency"`
}

func MustLoad() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("config path can not be empty")
	}

	if _, err := os.Stat(path); err != nil {
		panic("failed tp find config file: " + path)
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
