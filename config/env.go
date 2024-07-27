package config

import (
	"os"

	"github.com/spf13/viper"
)

var (
	PORT         string
	DATABASE_URL string
	BROKER_URL   string
)

func LoadEnv() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		PORT = os.Getenv("PORT")
		DATABASE_URL = os.Getenv("DATABASE_URL")
		BROKER_URL = os.Getenv("BROKER_URL")
		return
	}

	PORT = viper.GetString("PORT")
	DATABASE_URL = viper.GetString("DATABASE_URL")
	BROKER_URL = viper.GetString("BROKER_URL")
}
