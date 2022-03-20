package endpoints

import (
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type Endpoints struct {
	DB    *gorm.DB
	Oauth *oauth2.Config
}
