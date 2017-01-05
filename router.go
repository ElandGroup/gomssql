package main

import (
	"gomssql/service"
	"net/http"

	"github.com/labstack/echo"
)

func RegisterApi() {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello db")
	})

	api := e.Group("/api")
	v1 := api.Group("/v1")

	fruit := v1.Group("/fruit")
	fruit.GET("", service.Query)

}
