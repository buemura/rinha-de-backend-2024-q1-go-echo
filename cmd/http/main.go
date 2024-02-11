package main

import (
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/config"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/database"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/statement"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/transaction"
	"github.com/labstack/echo/v4"
)

func init() {
	config.LoadEnv()
	database.Connect()
}

func setupRoutes(e *echo.Echo) {
	statement.SetupRoutes(e)
	transaction.SetupRoutes(e)
}

func main() {
	e := echo.New()
	setupRoutes(e)
	host := "0.0.0.0:" + config.PORT
	e.Start(host)
}
