package routes

import "github.com/pavan2108/golang-url-shortner/internal/controller"

func InitiateUrlRoutes() {
	urlRouters := echoInstance.Group("/url")
	var urlController controller.UrlController
	urlRouters.POST("/shorten", urlController.CreateShortenUrl)
	urlRouters.GET("/shorten/:short_code", urlController.GetShortenUrl)
}
