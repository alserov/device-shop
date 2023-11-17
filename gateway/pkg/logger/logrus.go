package logger

import (
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"io"
	"os"
)

func New(path string) (*logrus.Logger, error) {
	_, err := os.Stat(path + "/log/logs.txt")
	if os.IsNotExist(err) {
		if err = os.MkdirAll(path+"/log", 0777); err != nil {
			return nil, err
		}

		f, err := os.Create(path + "/log/logs.txt")
		if err != nil {
			return nil, err
		}
		defer f.Close()
	}

	f, err := os.OpenFile(path+"/log/logs.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
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

	log.Info("logger initialized")

	return log, nil
}
