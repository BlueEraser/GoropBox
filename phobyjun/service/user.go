package service

import (
	"phobyjun/db"
	"phobyjun/model"
)

func CreateUser(userDto *model.User) (*model.User, error) {
	hashPassword, err := userDto.HashPassword(userDto.Password)
	if err != nil {
		return nil, err
	}
	user := model.User{
		Username: userDto.Username,
		Password: hashPassword,
		Email:    userDto.Email,
	}
	tx := db.Session.Create(&user)
	if err := tx.Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	tx := db.Session.Where("email = ?", email).First(&user)
	if err := tx.Error; err != nil {
		return nil, err
	}

	return &user, nil
}
