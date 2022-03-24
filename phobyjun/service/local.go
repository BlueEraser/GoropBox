package service

import (
	"io"
	"mime/multipart"
	"os"
	"phobyjun/model"
	"strings"
)

const (
	baseDir = "uploaded"
)

func UploadFileToLocal(fileDto *model.File, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	parsedDir := parseFileDir(fileDto)
	dstDir := strings.Join([]string{
		baseDir,
		parsedDir,
	}, "/")

	dst, err := os.Create(dstDir)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func parseFileDir(fileDto *model.File) string {
	fileDir := fileDto.FileDir
	fileName := fileDto.FileName

	dirWithName := fileDir + fileName
	parsedDir := strings.ReplaceAll(dirWithName, "/", "@")

	return parsedDir
}
