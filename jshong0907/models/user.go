package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(50);not null;uniqueIndex;comment:이메일" json:"email"`
	Password string `gorm:"type:varchar(50);not null;comment:비밀번호" json:"-"`
	NickName string `gorm:"type:varchar(10);not null;comment:닉네임" json:"nickName"`
	Files    []File `gorm:"constraint:OnDelete:CASCADE"`
}

func (user User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err == nil {
		user.Password = string(hashedPassword)
		return nil
	}
	return err
}
