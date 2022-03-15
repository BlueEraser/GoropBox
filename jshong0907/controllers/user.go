package controllers

import (
	"github.com/golang-jwt/jwt"
	"gorop-box/models"
	"gorop-box/services"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {
	params := make(map[string]string)
	c.Bind(&params)
	user := services.CreateUser(
		params["email"],
		params["password"],
		params["nickName"],
	)

	return c.JSON(http.StatusOK, user)
}

func SignIn(c echo.Context) error {
	params := make(map[string]string)
	c.Bind(&params)
	user, err := services.CheckPassword(params["email"], params["password"])

	if err != nil {
		return echo.ErrUnauthorized
	}

	jwtClaims := models.JwtClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
		user.Email,
		user.NickName,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func GetUserInfo(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtClaims)
	return c.JSON(http.StatusOK, claims)
}
