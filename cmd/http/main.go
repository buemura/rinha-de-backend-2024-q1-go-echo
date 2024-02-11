package main

import (
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/config"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/modules/statement"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/modules/transaction"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/shared/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	config.LoadEnv()
	database.Connect()
}

func main() {
	e := echo.New()
	setupServerMiddlewares(e)
	host := ":" + config.PORT
	e.Start(host)
}

func setupServerMiddlewares(app *echo.Echo) {
	app.Use(middleware.Recover())
	app.Use(middleware.Secure())
	statement.SetupRoutes(app)
	transaction.SetupRoutes(app)
}
