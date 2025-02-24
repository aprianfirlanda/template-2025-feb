package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()

	level, err := logrus.ParseLevel(viper.GetString("LOG_LEVEL"))
	if err != nil {
		level = logrus.InfoLevel
	}
	Logger.SetLevel(level)

	Logger.SetFormatter(&logrus.JSONFormatter{})
}
