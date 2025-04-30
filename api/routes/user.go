package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pavan2108/golang-echo-template/global"
)

var echoInstance *echo.Echo

func init() {
	echoInstance = global.ECHO_INSTANCE
}

func InitiateUserRoutes() {
	echoInstance.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
}
