package errors

type InvalidPaswordError struct{}

func (m *InvalidPaswordError) Error() string {
	return "Invalid Password Error!"
}
