package controllers

import (
	"github.com/labstack/echo/v4"
	"gorop-box/auth"
	"gorop-box/services"
	"net/http"
)

func UploadFile(c echo.Context) error {
	user, jwtErr := auth.GetUserByJwt(c)
	if jwtErr != nil {
		return c.String(http.StatusBadRequest, jwtErr.Error())
	}

	fileForm, formErr := c.FormFile("file")
	if formErr != nil {
		return c.String(http.StatusBadRequest, formErr.Error())
	}

	fileSrc, err := fileForm.Open()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	defer fileSrc.Close()

	file, uploadErr := services.CreateFile(*user, fileSrc, fileForm.Filename)
	if uploadErr != nil {
		return c.String(http.StatusBadRequest, uploadErr.Error())
	}

	return c.JSON(http.StatusOK, file)
}
