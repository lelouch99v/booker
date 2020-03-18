package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const port = "5010"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, booker!")
	})
	e.Logger.Fatal(e.Start(":" + port))
}
