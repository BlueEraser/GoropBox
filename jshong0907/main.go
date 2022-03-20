package main

import (
	"github.com/labstack/echo/v4"
	"gorop-box/routers"
)

func main() {
	e := echo.New()
	routers.Init(e)

	e.Logger.Fatal(e.Start(":1323"))
}
