package controller

import (
	"log"
	"net/http"
	"os"
	"phobyjun/model"
	"phobyjun/service"
	"phobyjun/session"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UploadFile(c echo.Context) error {
	fileDto := &model.File{
		FileName:    c.FormValue("fileName"),
		FileNameDir: c.FormValue("fileNameDir"),
	}

	formFile, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	sess := session.Get(c)
	userId := sess.Values["userid"]
	if userId == nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	aesKey, hmacKey, err := service.GetKeysFromUserById(userId.(uint))
	fileDto, err = service.UploadFileToLocal(fileDto, formFile, aesKey, hmacKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	file, err := service.CreateFile(fileDto, userId.(uint))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, file)
}

func DownloadFile(c echo.Context) error {
	fileIdParam := c.Param("fileId")
	fileId, err := strconv.ParseUint(fileIdParam, 16, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	fileDto, err := service.GetFileByFileId(uint(fileId))
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}

	sess := session.Get(c)
	userId := sess.Values["userid"]
	if userId == nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	aesKey, hmacKey, err := service.GetKeysFromUserById(userId.(uint))
	fileDir, err := service.DownloadFileFromLocal(fileDto, aesKey, hmacKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	defer func(fileDir []byte) {
		err := os.Remove(string(fileDir))
		if err != nil {
			log.Fatal(err)
		}
	}(fileDir)

	return c.File(string(fileDir))
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
