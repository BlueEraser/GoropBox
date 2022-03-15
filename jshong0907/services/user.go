package services

import (
	"gorop-box/errors"
	"gorop-box/models"
)

func CreateUser(email, password, nickName string) models.User {
	user := models.User{Email: email, NickName: nickName}
	user.SetPassword(password)
	DB.Create(&user)
	return user
}

func GetUserByEmail(email string) models.User {
	var user models.User
	DB.Where("email = ?", email).Find(&user)
	return user
}

func CheckPassword(email, password string) (models.User, error) {
	var user models.User
	DB.Where("email = ?", email).Find(&user)
	if user.CheckPassword(password) {
		return user, nil
	}
	return user, &errors.InvalidPasswordError{}
}
