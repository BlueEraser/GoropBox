package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorop-box/auth"
	"gorop-box/controllers"
)

func setUserRouter(r *echo.Group) {
	config := middleware.JWTConfig{
		Claims:     &auth.JwtClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("/user", controllers.GetUserInfo)
	r.GET("/file", controllers.UploadFile)
}
