package models

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Email     string
	FileName  string
	Path      string
	Extension string
	FileSize  int64
}
