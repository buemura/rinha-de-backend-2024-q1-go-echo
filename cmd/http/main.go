package main

import (
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/config"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/modules/statement"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/modules/transaction"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/shared/database"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/shared/helper"
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
	app.JSONSerializer = helper.CustomJsonSerializer{Provider: "sonic"}
	app.Use(middleware.Recover())
	app.Use(middleware.Secure())
	statement.SetupRoutes(app)
	transaction.SetupRoutes(app)
}
