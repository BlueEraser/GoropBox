package service

import (
	"encoding/base64"
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
	src, err := os.Open(file.Filename)
	if err != nil {
		return err
	}
	defer src.Close()

	encryptedName := encryptFileNameDir(fileDto)
	dstDir := strings.Join([]string{
		baseDir,
		encryptedName,
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

func encryptFileNameDir(fileDto *model.File) string {
	fileNameDir := fileDto.FileNameDir

	return base64.StdEncoding.EncodeToString([]byte(fileNameDir))
}
