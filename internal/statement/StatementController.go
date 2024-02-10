package statement

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/customer"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/clientes/:clienteId/extrato", getStatement)
}

func getStatement(c echo.Context) error {
	customerIdStr := c.Param("clienteId")
	customerId, err := strconv.Atoi(customerIdStr)
	if err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	stt, err := GetStatement(customerId)

	if err != nil {
		switch {
		case errors.Is(err, customer.ErrCustomerNotFound):
			return c.NoContent(http.StatusNotFound)
		default:
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	return c.JSON(http.StatusOK, stt)
}
