package services

import (
	"gorop-box/errors"
	"gorop-box/models"
)

func CreateUser(email, password, nickName string) (*models.User, error) {
	user := models.User{Email: email, NickName: nickName}
	if email == "" || password == "" || nickName == "" {
		return nil, &errors.ValidationError{}
	}
	err := user.SetPassword(password)
	if err != nil {
		return nil, err
	}
	DB.Create(&user)
	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	DB.Where("email = ?", email).Find(&user)
	return &user, nil
}

func CheckPassword(email, password string) (*models.User, error) {
	var user models.User
	DB.Where("email = ?", email).Find(&user)
	if user.CheckPassword(password) != nil {
		return &user, nil
	}
	return &user, &errors.InvalidPasswordError{}
}
