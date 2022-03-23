package router

import (
	"github.com/labstack/echo/v4"
	"phobyjun/controller"
)

func setFileController(e *echo.Echo) {
	e.POST(APIFile, controller.UploadFile)
}
