package routers

import (
	"github.com/labstack/echo/v4"
	"gorop-box/controllers"
)

func setAuthRouter(r *echo.Group) {
	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.SignIn)
}
