package errors

type InvalidPasswordError struct{}

func (m *InvalidPasswordError) Error() string {
	return "Invalid Password Error!"
}
