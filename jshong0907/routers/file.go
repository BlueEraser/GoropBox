package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorop-box/auth"
	"gorop-box/controllers"
)

func setFileRouter(r *echo.Group) {
	config := middleware.JWTConfig{
		Claims:     &auth.JwtClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.POST("", controllers.UploadFile)
	r.GET("/:file", controllers.GetFile)
	r.DELETE("/:file", controllers.DeleteFile)
	r.DELETE("", controllers.DeleteAllFiles)
}
