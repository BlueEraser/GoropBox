package services

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var awsSession *session.Session
var s3Uploader *s3manager.Uploader

func init() {
	var sessionErr error
	awsSession, sessionErr = session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	if sessionErr != nil {
		log.Fatal("Error Loading AWS session")
	}

	s3Uploader = s3manager.NewUploader(awsSession)
}
