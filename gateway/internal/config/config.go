package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env      string        `yaml:"env"`
	Port     int           `yaml:"port"`
	Timeout  time.Duration `yaml:"timeout"`
	Cache    Cache         `yaml:"cache"`
	Services Services      `yaml:"services"`
}

type Services struct {
	User   Service `yaml:"user"`
	Device Service `yaml:"device"`
	Order  Service `yaml:"order"`
	Coll   Service `yaml:"coll"`
}

type Service struct {
	Addr string `yaml:"addr"`
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
