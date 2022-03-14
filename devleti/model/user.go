package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Files    []File `gorm:"foreignKey:UploadUser"`
}
