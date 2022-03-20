package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func UploadFile(file *os.File) error {
	Uploader = s3manager.NewUploader(Session)
	_, err := Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(file.Name()),
		Body:   file,
	})

	if err != nil {
		return err
	}
	return nil
}
