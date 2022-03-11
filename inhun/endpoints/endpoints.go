package endpoints

import (
	"gorm.io/gorm"
)

type Endpoints struct {
	DB *gorm.DB
}
