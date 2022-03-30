package main

import (
	"phobyjun/router"
	"phobyjun/session"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	session.Init(e)
	router.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}
