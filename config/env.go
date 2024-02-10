package config

import "github.com/spf13/viper"

var (
	PORT         string
	DATABASE_URL string
)

func LoadEnv() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic("Failed to load environment variables")
	}

	PORT = viper.GetString("PORT")
	DATABASE_URL = viper.GetString("DATABASE_URL")
}
