package model

import "time"

type File struct {
	ID            uint      `gorm:"primary_key" json:"id"`
	FileName      string    `json:"fileName"`
	FileNameDir   string    `gorm:"unique" json:"fileNameDir"`
	EncryptedName string    `json:"-"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"UpdatedAt"`
	UserID        uint
}
