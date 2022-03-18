package validator

import "github.com/go-playground/validator/v10"

type UserValidator struct {
	validator *validator.Validate
}

func (uv UserValidator) Validate(u interface{}) error {
	if err := uv.validator.Struct(u); err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}
