package errors

type InvalidPasswordError struct{}

func (m *InvalidPasswordError) Error() string {
	return "Invalid Password Error!"
}

type ValidationError struct{}

func (m *ValidationError) Error() string {
	return "Validation Error!"
}
