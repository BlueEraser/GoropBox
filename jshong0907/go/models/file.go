package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	UserID      uint   `json:"user"`
	Path        string `gorm:"type:varchar(255);not null;comment:경로" json:"path"`
	IsDirectory bool   `gorm:"not null;comment:폴더여부" json:"is_directory"`
	AbsoluteUrl string `gorm:"-:all" json:"url"`
}
