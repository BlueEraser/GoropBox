package services

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gorm.io/gorm"
	"gorop-box/box_errors"
	"gorop-box/models"
	"io"
	"os"
)

func uploadFile(file io.Reader, fileName string) error {
	s3Uploader = s3manager.NewUploader(awsSession)
	_, err := s3Uploader.Upload(&s3manager.UploadInput{
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
	db.Create(&file)
	return &file, nil
}

func GetFile(user models.User, fileName string) (*models.File, error) {
	var file models.File
	result := db.Where("user_id = ? AND path = ?", user.ID, fileName).Take(&file)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &box_errors.ValidationError{ErrorMessage: "등록되지 않은 파일입니다."}
	}
	return &file, nil
}

func DeleteFile(user models.User, fileName string) error {
	file, err := GetFile(user, fileName)
	if err != nil {
		return err
	}
	db.Delete(&file)
	return nil
}
