package services

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
)

var Session *session.Session
var Uploader *s3manager.Uploader

func init() {
	var sessionErr error
	Session, sessionErr = session.NewSession()
	if sessionErr != nil {
		log.Fatal("Error Loading AWS session")
	}

	Uploader = s3manager.NewUploader(Session)
}
