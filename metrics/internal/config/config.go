package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env    string `yaml:"env"`
	Server Server `yaml:"server"`
	Broker Broker `yaml:"broker"`
}

type Server struct {
	Port int `yaml:"port"`
}

type Broker struct {
	Addr   string `yaml:"addr"`
	Topics Topics `yaml:"topics"`
}

type Topics struct {
	Request string `yaml:"request"`
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
