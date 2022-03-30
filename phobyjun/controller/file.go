package controller

import (
	"net/http"
	"phobyjun/model"
	"phobyjun/service"
	"phobyjun/session"

	"github.com/labstack/echo/v4"
)

func UploadFile(c echo.Context) error {
	fileDto := &model.File{
		FileName:    c.FormValue("fileName"),
		FileNameDir: c.FormValue("fileDir"),
	}

	formFile, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := service.UploadFileToLocal(fileDto, formFile); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	sess := session.Get(c)
	userId := sess.Values["userid"]
	if userId == nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	file, err := service.CreateFile(fileDto, userId.(uint))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, file)
}

func ListFiles(c echo.Context) error {
	sess := session.Get(c)
	userId := sess.Values["userid"]
	if userId == nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	files, err := service.ListFilesByUserId(userId.(uint))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, files)
}
