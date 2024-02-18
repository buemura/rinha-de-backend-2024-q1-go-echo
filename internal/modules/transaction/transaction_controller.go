package transaction

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/shared/helper"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.POST("/clientes/:clienteId/transacoes", createTransaction)
}

func createTransaction(c echo.Context) error {
	customerIdStr := c.Param("clienteId")
	customerId, err := strconv.Atoi(customerIdStr)
	if err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}
	body := new(CreateTransactionRequest)
	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}
	if err := validate(body); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}
	trx, err := CreateTransaction(customerId, body)
	if err != nil {
		return helper.HandleHttpError(c, err)
	}
	return c.JSON(http.StatusOK, trx)
}

func validate(body *CreateTransactionRequest) error {
	if body.Amount < 1 {
		return errors.New("invalid amount")
	}
	if len(body.Description) < 1 || len(body.Description) > 10{
		return errors.New("invalid description")
	}
	if body.Type != "c" && body.Type != "d" {
		return errors.New("invalid type")
	}
	return nil
}