package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"phobyjun/model"
	"phobyjun/model/validator"
	"phobyjun/service"
)

func SignUp(c echo.Context) error {
	userDto := &model.User{}

	if err := c.Bind(userDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fmt.Println(userDto)
	if err := validator.UserValidate(userDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := service.CreateUser(userDto)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
