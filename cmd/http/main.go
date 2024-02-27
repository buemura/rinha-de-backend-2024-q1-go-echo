package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/config"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/modules/statement"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/modules/transaction"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/shared/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setupServerMiddlewares(app *echo.Echo) {
	app.Use(middleware.Recover())
	app.Use(middleware.Secure())
	statement.SetupRoutes(app)
	transaction.SetupRoutes(app)
}

func init() {
	config.LoadEnv()
	database.Connect()
}

func main() {
	e := echo.New()
	setupServerMiddlewares(e)
	host := ":" + config.PORT

	go func ()  {
		if err := e.Start(host); err != nil && http.ErrServerClosed != err {
			panic(err)
		  } 
	}()

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	<-stop
	
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	fmt.Println("Stopping...")
	
	if err := e.Shutdown(ctx); err != nil {
	  panic(err)
	}
	fmt.Println("Server stopped")	
}