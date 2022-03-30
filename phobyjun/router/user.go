package router

import (
	"phobyjun/controller"

	"github.com/labstack/echo/v4"
)

func setAuthController(e *echo.Echo) {
	e.POST(APIAuthSignup, controller.SignUp)
	e.POST(APIAuth, controller.LogIn)
	e.DELETE(APIAuth, controller.LogOut)

	e.GET(APIAuth, controller.GetUserInfo)
}
