package services

import (
	"gorop-box/models"
)

func CreateUser(email, password, nickName string) models.User {
	user := models.User{Email: email, NickName: nickName}
	user.SetPassword(password)
	DB.Create(&user)
	return user
}
