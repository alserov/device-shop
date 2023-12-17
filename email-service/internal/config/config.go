package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env   string `yaml:"env"`
	Email struct {
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Email    string `yaml:"email"`
	} `yaml:"email"`
	Broker struct {
		Addr   string `yaml:"addr"`
		Topics struct {
			AuthTopic  string `yaml:"authTopic"`
			OrderTopic string `yaml:"orderTopic"`
		} `yaml:"topics"`
	} `yaml:"broker"`
}

func MustLoadEmail() *Config {
	path := fetchEmailConfigPath()

	if _, err := os.Stat(path); err != nil {
		panic("config file not found: " + path)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchEmailConfigPath() string {
	var path string

	flag.StringVar(&path, "c", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
