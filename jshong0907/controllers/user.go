package controllers

import (
	"net/http"

	"gorop-box/services"

	"github.com/labstack/echo/v4"
)

func Signup(c echo.Context) error {
	params := make(map[string]string)
	c.Bind(&params)
	user := services.CreateUser(
		params["email"],
		params["password"],
		params["nickName"],
	)

	return c.JSON(http.StatusOK, user)
}
