package service

import (
	"encoding/base64"
	"io"
	"log"
	"mime/multipart"
	"os"
	"phobyjun/model"
	"strings"
)

const (
	baseDir = "uploaded"
)

func UploadFileToLocal(fileDto *model.File, file *multipart.FileHeader) (*model.File, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(src)

	encryptedName := encryptFileNameDir(fileDto)
	fileDto.EncryptedName = encryptedName

	dstDir := strings.Join([]string{
		baseDir,
		encryptedName,
	}, "/")

	dst, err := os.Create(dstDir)
	if err != nil {
		return nil, err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dst)

	if _, err := io.Copy(dst, src); err != nil {
		return nil, err
	}

	return fileDto, nil
}

func encryptFileNameDir(fileDto *model.File) string {
	fileNameDir := fileDto.FileNameDir

	return base64.StdEncoding.EncodeToString([]byte(fileNameDir))
}
