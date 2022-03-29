package services

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var awsSession *session.Session
var s3Uploader *s3manager.Uploader
var s3Session *s3.S3
var s3Url string

func init() {
	var sessionErr error
	awsSession, sessionErr = session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	if sessionErr != nil {
		log.Fatal("Error Loading AWS session")
	}

	s3Uploader = s3manager.NewUploader(awsSession)
	s3Session = s3.New(awsSession)
	s3Url = fmt.Sprintf(
		"https://%s.s3-%s.amazonaws.com/",
		os.Getenv("AWS_BUCKET"),
		os.Getenv("AWS_REGION"),
	)
}
