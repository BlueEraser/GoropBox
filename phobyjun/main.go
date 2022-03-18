package main

import (
	"github.com/labstack/echo/v4"
	"phobyjun/config"
	"phobyjun/controller"
	"phobyjun/db"
)

func main() {
	config.Init()
	db.Init()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello Echo!")
	})
	e.POST("/user", controller.SignUp)
	e.Logger.Fatal(e.Start(":8080"))
}
