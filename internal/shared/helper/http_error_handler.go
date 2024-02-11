package helper

import (
	"errors"
	"net/http"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo/internal/modules/customer"
	"github.com/labstack/echo/v4"
)

func HandleHttpError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, customer.ErrCustomerNotFound):
		return c.NoContent(http.StatusNotFound)
	case errors.Is(err, customer.ErrCustomerNoLimit):
		return c.NoContent(http.StatusUnprocessableEntity)
	default:
		return c.NoContent(http.StatusInternalServerError)
	}
}
