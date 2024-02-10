package main

import (
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/statement"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var (
	PORT         string
	DATABASE_URL string
)

func init() {
	loadEnvVariables()
	database.Connect()
}

func loadEnvVariables() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic("Failed to load environment variables")
	}

	PORT = viper.GetString("PORT")
	DATABASE_URL = viper.GetString("DATABASE_URL")
}

func setupRoutes(e *echo.Echo) {
	statement.SetupRoutes(e)
}

func main() {
	e := echo.New()
	setupRoutes(e)
	host := "127.0.0.1:" + PORT
	e.Start(host)
}
