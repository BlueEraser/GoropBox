package model

import "time"

type File struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	FileName  string    `json:"fileName"`
	FileDir   string    `json:"fileDir"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"UpdatedAt"`
	UserID    uint
}
