package main

import "github.com/labstack/echo"

var (
	e = echo.New()
)

func main() {
	RegisterApi()

	e.Start(":1116")
}
