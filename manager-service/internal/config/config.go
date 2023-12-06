package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type EmailConfig struct {
	Env   string `yaml:"env"`
	Email struct {
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Email    string `yaml:"email"`
	} `yaml:"email"`
	Topics struct {
		Email struct {
			AuthTopic  string `yaml:"authTopic"`
			OrderTopic string `yaml:"orderTopic"`
		} `yaml:"email"`
	} `yaml:"topics"`
	BrokerAddr string `yaml:"brokerAddr"`
}

type TxConfig struct {
	Env    string `yaml:"env"`
	Topics struct {
		Tx struct {
			Topic    string `yaml:"topic"`
			Services struct {
				UserService       string `yaml:"userService"`
				DeviceService     string `yaml:"deviceService"`
				CollectionService string `yaml:"collectionService"`
			} `yaml:"services"`
		} `yaml:"tx"`
	} `yaml:"topics"`
	BrokerAddr string `yaml:"brokerAddr"`
}

func MustLoadTransaction() *TxConfig {
	path := fetchEmailConfigPath()

	if _, err := os.Stat(path); err != nil {
		panic("config file not found: " + path)
	}

	var cfg TxConfig
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
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
