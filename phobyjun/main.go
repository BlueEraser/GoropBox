package main

import (
	"github.com/labstack/echo/v4"
	"phobyjun/db"
)

func main() {
	db.Init()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello Echo!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
