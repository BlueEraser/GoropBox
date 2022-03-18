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
	db.Session.Create(&user)
	return &user, nil
}
