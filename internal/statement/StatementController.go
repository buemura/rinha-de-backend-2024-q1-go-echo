package statement

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/clientes/:clienteId/extrato", getStatement)
}

func getStatement(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"message": "Extratoooo",
	})
}
