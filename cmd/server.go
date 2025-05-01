package cmd

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pavan2108/golang-url-shortner/api/routes"
	"github.com/pavan2108/golang-url-shortner/configs/environment"
	"github.com/pavan2108/golang-url-shortner/global"
)

var echoInstance *echo.Echo = global.ECHO_INSTANCE

func AddMiddleware() {
	echoInstance.Use(middleware.Logger())
}

func init() {
	AddMiddleware()
	routes.InitiateUserRoutes()
	routes.InitiateUrlRoutes()
}

func StartServer() {
	var port string = ":" + strconv.Itoa(environment.PORT)
	fmt.Println(port)
	echoInstance.Start(port)
}
