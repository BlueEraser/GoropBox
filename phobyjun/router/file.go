package router

import (
	"phobyjun/controller"

	"github.com/labstack/echo/v4"
)

func setFileController(e *echo.Echo) {
	e.POST(APIFile, controller.UploadFile)
	e.GET(APIFile, controller.ListFiles)
}
