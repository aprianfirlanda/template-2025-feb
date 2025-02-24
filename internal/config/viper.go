package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig(fileName string) {
	viper.AutomaticEnv()

	viper.SetConfigFile(fileName)
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("There is no env file. It is a problem if you run in local mode. If it is not, then is ok.")
	}
}
