package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"phobyjun/db"
	"phobyjun/model"
	"phobyjun/model/validator"
	"phobyjun/service"
	"phobyjun/session"
)

func SignUp(c echo.Context) error {
	userDto := &model.User{}
	if err := c.Bind(userDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := validator.UserValidate(userDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := service.CreateUser(userDto)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func LogIn(c echo.Context) error {
	userDto := &model.User{}
	if err := c.Bind(userDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	paramEmail := userDto.Email
	paramPassword := userDto.Password

	var user model.User
	db.Session.Where("email = ?", paramEmail).First(&user)

	if err := user.CheckPassword(paramPassword); err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}

	if err := session.Save(c, paramEmail); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
