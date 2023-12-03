package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type EmailConfig struct {
	Email Email  `yaml:"email"`
	Kafka Kafka  `yaml:"kafka"`
	Env   string `yaml:"env"`
}

type Kafka struct {
	Topics     Topic  `yaml:"topics"`
	BrokerAddr string `yaml:"brokerAddr"`
}

type Topic struct {
	Auth  string `yaml:"auth"`
	Order string `yaml:"order"`
}

type Email struct {
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Email    string `yaml:"email"`
}

func MustLoadEmail() *EmailConfig {
	path := fetchEmailConfigPath()

	if _, err := os.Stat(path); err != nil {
		panic("config file not found: " + path)
	}

	var cfg EmailConfig
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
