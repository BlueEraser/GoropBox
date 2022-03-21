package main

import (
	"github.com/labstack/echo/v4"
	"phobyjun/router"
	"phobyjun/session"
)

func main() {
	e := echo.New()

	session.Init(e)
	router.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}
