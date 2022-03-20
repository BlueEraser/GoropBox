package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gorop-box/models"
	"io"
	"os"
)

func uploadFile(file io.Reader, fileName string) error {
	Uploader = s3manager.NewUploader(Session)
	_, err := Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(fileName),
		Body:   file,
	})

	if err != nil {
		return err
	}
	return nil
}

func CreateFile(user models.User, fileReader io.Reader, fileName string) (*models.File, error) {
	file := models.File{
		UserID:      user.ID,
		Path:        fileName,
		IsDirectory: false,
	}
	err := uploadFile(fileReader, fileName)
	if err != nil {
		return nil, err
	}
	DB.Create(&file)
	return &file, nil
}
