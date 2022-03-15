package main

import (
	"net/http"
	"time"

	"gorop-box/controllers"
	"gorop-box/services"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type jwtClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func main() {
	e := echo.New()
	e.POST("/signup", controllers.Signup)
	e.POST("/signin", func(c echo.Context) error {
		params := make(map[string]string)
		c.Bind(&params)
		user, err := services.Signin(params["email"], params["password"])

		if err != nil {
			return echo.ErrUnauthorized
		}

		claims := jwtClaims{
			user.NickName,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	})

	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &jwtClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":1323"))
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtClaims)
	email := claims.Email
	return c.String(http.StatusOK, "Welcome "+email+"!")
}
