package services

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"os"

	"gorop-box/box_errors"
	"gorop-box/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gorm.io/gorm"
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
	file.AbsoluteUrl = fmt.Sprintf("%s%s", s3Url, file.Path)
	return &file, nil
}

func DeleteFile(user models.User, fileName string) error {
	file, err := GetFile(user, fileName)
	if err != nil {
		return err
	}

	db.Delete(&file)
	_, err = s3Session.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return err
	}

	err = s3Session.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return err
	}
	return nil
}
