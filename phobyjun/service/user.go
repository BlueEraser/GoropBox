package service

import (
	"phobyjun/db"
	"phobyjun/model"
)

func CreateUser(username string, password string, email string) (*model.User, error) {
	hashPassword, err := model.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := model.User{Username: username, Password: hashPassword, Email: email}
	db.Session.Create(&user)
	return &user, nil
}
