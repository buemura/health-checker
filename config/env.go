package config

import (
	"os"

	"github.com/spf13/viper"
)

var (
	PORT         string
	DATABASE_URL string
	BROKER_URL   string

	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASS     string
	DB_DATABASE string
)

func LoadEnv() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		PORT = os.Getenv("PORT")
		DATABASE_URL = os.Getenv("DATABASE_URL")
		BROKER_URL = os.Getenv("BROKER_URL")

		DB_HOST = os.Getenv("DB_HOST")
		DB_PORT = os.Getenv("DB_PORT")
		DB_USER = os.Getenv("DB_USER")
		DB_PASS = os.Getenv("DB_PASS")
		DB_DATABASE = os.Getenv("DB_DATABASE")
		return
	}

	PORT = viper.GetString("PORT")
	DATABASE_URL = viper.GetString("DATABASE_URL")
	BROKER_URL = viper.GetString("BROKER_URL")

	DB_HOST = viper.GetString("DB_HOST")
	DB_PORT = viper.GetString("DB_PORT")
	DB_USER = viper.GetString("DB_USER")
	DB_PASS = viper.GetString("DB_PASS")
	DB_DATABASE = viper.GetString("DB_DATABASE")
}
