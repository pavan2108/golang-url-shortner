package global

import (
	"github.com/labstack/echo/v4"
)

var ECHO_INSTANCE *echo.Echo = echo.New()

type (
	ErrorObject struct {
		Message string
	}

	ResponseStructure struct {
		StatusCode int         `json:"statusCode"`
		Message    string      `json:"message"`
		Error      ErrorObject `json:"error"`
		Data       any         `json:"data"`
	}
)
