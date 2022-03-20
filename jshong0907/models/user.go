package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(50);not null;uniqueIndex;comment:이메일" json:"email"`
	Password string `gorm:"type:varchar(255);not null;comment:비밀번호" json:"-"`
	NickName string `gorm:"type:varchar(10);not null;comment:닉네임" json:"nickName"`
	Files    []File `gorm:"constraint:OnDelete:CASCADE"`
}

// CheckPassword 는 User 의 복호화된 Password 와 파라미터 password 를 비교합니다.
// password 와 복호화된 Password 가 동일하다면 nil 을, 다르다면 error 를 반환합니다.
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
