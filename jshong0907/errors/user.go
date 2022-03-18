package errors

import "fmt"

type InvalidPasswordError struct{}

func (m *InvalidPasswordError) Error() string {
	return "Invalid Password Error!"
}

type ValidationError struct {
	ErrorMessage string
}

func (m *ValidationError) Error() string {
	return fmt.Sprintf("Validation Error: %s", m.ErrorMessage)
}
