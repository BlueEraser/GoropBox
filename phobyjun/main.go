package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"phobyjun/config"
	"phobyjun/db"
)

func main() {
	config.Init()
	db.Init()

	fmt.Printf("%+v\n", config.Cfg)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello Echo!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
