package logger

import (
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"io"
	"os"
)

func New() (*logrus.Logger, error) {
	f, err := os.OpenFile("gateway/log/logs.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	log := &logrus.Logger{
		Out: io.MultiWriter(f, os.Stdout),
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
	}

	return log, nil
}
