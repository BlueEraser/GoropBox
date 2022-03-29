package controllers

import (
	"fmt"
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

	fileName := fmt.Sprintf("%s/%s", c.FormValue("dir"), fileForm.Filename)
	file, uploadErr := services.CreateFile(*user, fileSrc, fileName)
	if uploadErr != nil {
		return c.String(http.StatusBadRequest, uploadErr.Error())
	}

	return c.JSON(http.StatusOK, file)
}

func GetFile(c echo.Context) error {
	user, jwtErr := auth.GetUserByJwt(c)
	if jwtErr != nil {
		return c.String(http.StatusBadRequest, jwtErr.Error())
	}

	fileName := c.Param("file")
	file, err := services.GetFile(*user, fileName)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, file)
}

func DeleteFile(c echo.Context) error {
	user, jwtErr := auth.GetUserByJwt(c)
	if jwtErr != nil {
		return c.String(http.StatusBadRequest, jwtErr.Error())
	}

	fileName := c.Param("file")
	err := services.DeleteFile(*user, fileName)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusBadRequest, "파일이 성공적으로 삭제되었습니다.")
}
