package config

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestInitLogger(t *testing.T) {
	originalLogLevel := os.Getenv("LOG_LEVEL")
	defer os.Setenv("LOG_LEVEL", originalLogLevel)

	testCases := []struct {
		name          string
		logLevel      string
		expectedLevel logrus.Level
	}{
		{"Default Level (INFO)", "", logrus.InfoLevel},
		{"Debug Level", "debug", logrus.DebugLevel},
		{"Warn Level", "warn", logrus.WarnLevel},
		{"Error Level", "error", logrus.ErrorLevel},
		{"Invalid Level (Fallback to INFO)", "invalid", logrus.InfoLevel},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("LOG_LEVEL", tc.logLevel)
			viper.AutomaticEnv()
			InitLogger()

			assert.Equal(t, tc.expectedLevel, Logger.GetLevel())
		})
	}
}
