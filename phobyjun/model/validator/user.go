package validator

import "github.com/go-playground/validator/v10"

var UserValidator *validator.Validate

func init() {
	UserValidator = validator.New()
}

func UserValidate(user interface{}) error {
	if err := UserValidator.Struct(user); err != nil {
		return err.(validator.ValidationErrors)
	}

	return nil
}
