package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorop-box/controllers"
	"gorop-box/models"
)

func main() {
	e := echo.New()
	e.POST("/signup", controllers.SignUp)
	e.POST("/signin", controllers.SignIn)

	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &models.JwtClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("/user", controllers.GetUserInfo)

	e.Logger.Fatal(e.Start(":1323"))
}
