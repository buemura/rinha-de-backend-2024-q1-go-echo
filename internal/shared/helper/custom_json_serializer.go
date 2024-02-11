package helper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/labstack/echo/v4"
)

type CustomJsonSerializer struct {
	Provider string
}

func (d CustomJsonSerializer) Serialize(c echo.Context, i interface{}, indent string) error {
	switch d.Provider {
	case "sonic":
		return sonicSerializer(c, i, indent)
	default:
		return jsonSerializer(c, i, indent)
	}
}

func (d CustomJsonSerializer) Deserialize(c echo.Context, i interface{}) error {
	switch d.Provider {
	case "sonic":
		return sonicDeserializer(c, i)
	default:
		return jsonDeserializer(c, i)
	}
}

func sonicSerializer(c echo.Context, i interface{}, indent string) error {
	enc := sonic.ConfigDefault.NewEncoder(c.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(i)
}

func sonicDeserializer(c echo.Context, i interface{}) error {
	err := sonic.ConfigDefault.NewDecoder(c.Request().Body).Decode(i)
	if ute, ok := err.(*json.UnmarshalTypeError); ok {
		return c.String(http.StatusUnprocessableEntity, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset))
	} else if se, ok := err.(*json.SyntaxError); ok {
		return c.String(http.StatusUnprocessableEntity, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error()))
	}
	return err
}

func jsonSerializer(c echo.Context, i interface{}, indent string) error {
	enc := json.NewEncoder(c.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(i)
}

func jsonDeserializer(c echo.Context, i interface{}) error {
	err := json.NewDecoder(c.Request().Body).Decode(i)
	if ute, ok := err.(*json.UnmarshalTypeError); ok {
		return c.String(http.StatusUnprocessableEntity, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset))
	} else if se, ok := err.(*json.SyntaxError); ok {
		return c.String(http.StatusUnprocessableEntity, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error()))
	}
	return err
}
