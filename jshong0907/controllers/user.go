package controllers

import (
	"net/http"
	"time"

	"gorop-box/auth"
	"gorop-box/services"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {
	params := make(map[string]string)
	if bindErr := c.Bind(&params); bindErr != nil {
		return bindErr
	}
	user, err := services.CreateUser(
		params["email"],
		params["password"],
		params["nickName"],
	)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func SignIn(c echo.Context) error {
	params := make(map[string]string)
	if bindErr := c.Bind(&params); bindErr != nil {
		return bindErr
	}
	user, err := services.CheckPassword(params["email"], params["password"])

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	jwtClaims := auth.JwtClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
		user.Email,
		user.NickName,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func GetUserInfo(c echo.Context) error {
	userJwt := c.Get("user").(*jwt.Token)
	claims := userJwt.Claims.(*auth.JwtClaims)
	user, _ := services.GetUserByEmail(claims.Email)
	return c.JSON(http.StatusOK, user)
}
