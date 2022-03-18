package box_errors

import "fmt"

type InvalidPasswordError struct{}

func (m *InvalidPasswordError) Error() string {
	return "비밀번호가 일치하지 않습니다."
}

type ValidationError struct {
	ErrorMessage string
}

func (m *ValidationError) Error() string {
	return fmt.Sprintf("Validation Error: %s", m.ErrorMessage)
}
