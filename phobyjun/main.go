package main

import (
	"github.com/labstack/echo/v4"
	"phobyjun/router"
)

func main() {
	e := echo.New()

	router.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}
