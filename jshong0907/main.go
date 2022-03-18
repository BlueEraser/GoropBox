package main

import (
	"gorop-box/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routers.Init(e)

	e.Logger.Fatal(e.Start(":1323"))
}
