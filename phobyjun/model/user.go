package model

import (
	"crypto/rand"

	scrypt "github.com/elithrar/simple-scrypt"
)

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `gorm:"unique" json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `gorm:"unique" json:"email" validate:"required,email"`
	AesKey   []byte `json:"-"`
	HmacKey  []byte `json:"-"`
	Files    []File
}

func (user User) HashPassword(password string) (string, error) {
	hash, err := scrypt.GenerateFromPassword([]byte(password), scrypt.DefaultParams)
	return string(hash), err
}

func (user User) CheckPassword(password string) error {
	return scrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (user User) Generate32ByteKey() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}
