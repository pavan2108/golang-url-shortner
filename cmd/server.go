package cmd

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pavan2108/golang-echo-template/api/routes"
	"github.com/pavan2108/golang-echo-template/configs/environment"
	"github.com/pavan2108/golang-echo-template/global"
)

var echoInstance *echo.Echo = global.ECHO_INSTANCE

func init() {
	routes.InitiateUserRoutes()
}

func StartServer() {
	var port string = ":" + strconv.Itoa(environment.PORT)
	fmt.Println(port)
	echoInstance.Start(port)
}
