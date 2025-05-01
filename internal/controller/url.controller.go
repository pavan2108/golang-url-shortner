package controller

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/pavan2108/golang-url-shortner/global"
	"github.com/pavan2108/golang-url-shortner/internal/url"
)

type UrlController struct {
}

var urlService url.UrlService

func (u *UrlController) CreateShortenUrl(c echo.Context) error {
	urlRequestBody := new(url.UrlShortenRequestType)
	err := c.Bind(urlRequestBody)
	if err != nil {
		return c.JSON(400, global.ResponseStructure{
			StatusCode: 400,
			Error: global.ErrorObject{
				Message: err.Error(),
			},
		})
	}
	var shortenUrl string = urlService.Shorten(urlRequestBody.Url)
	return c.JSON(200, global.ResponseStructure{
		StatusCode: 200,
		Message:    "Succes",
		Data: map[string]any{
			"url": shortenUrl,
		},
	})
}

func (u *UrlController) GetShortenUrl(c echo.Context) error {
	var shortCode string = c.Param("short_code")
	var originalUrl string = urlService.GetUrl(shortCode)
	if len(originalUrl) == 0 {
		return c.JSON(404, global.ResponseStructure{
			StatusCode: 404,
			Message:    "Not Found",
			Error: global.ErrorObject{
				Message: fmt.Sprintf("No Record found with error url : %s", shortCode),
			},
		})
	}
	return c.Redirect(301, originalUrl)

}
