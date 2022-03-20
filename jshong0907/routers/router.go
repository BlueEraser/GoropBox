package routers

import "github.com/labstack/echo/v4"

func Init(e *echo.Echo) {
	authGroup := e.Group("/auth")
	setAuthRouter(authGroup)

	userGroup := e.Group("/user")
	setUserRouter(userGroup)

	fileGroup := e.Group("file")
	setFileRouter(fileGroup)
}
