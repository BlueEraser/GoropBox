package endpoints

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type Endpoints struct {
	DB       *gorm.DB
	Oauth    *oauth2.Config
	S3Client *s3.Client
}
