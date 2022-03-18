package services

import (
	"gorop-box/errors"
	"gorop-box/models"
	"regexp"
)

func CreateUser(email, password, nickName string) (*models.User, error) {
	user := models.User{Email: email, NickName: nickName}
	if email == "" || password == "" || nickName == "" {
		return nil, &errors.ValidationError{ErrorMessage: "이메일, 패스워드, 닉네임이 모두 입력되어야합니다."}
	}
	matched, regrexErr := regexp.MatchString("[a-zA-Z0-9]@[a-zA-Z0-9].[a-zA-Z]", email)
	if regrexErr != nil {
		return nil, regrexErr
	}
	if !matched {
		return nil, &errors.ValidationError{ErrorMessage: "이메일이 올바르지 않은 형태입니다."}
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
