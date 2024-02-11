package transaction

import (
	"net/http"
	"strconv"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/utils"
	"github.com/go-playground/validator/v10"
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
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}
	if body.Tipo != "c" && body.Tipo != "d" {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	trx, err := CreateTransaction(customerId, body)
	if err != nil {
		return utils.HandleHttpError(c, err)
	}

	return c.JSON(http.StatusOK, trx)
}
