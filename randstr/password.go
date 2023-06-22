package randstr

// NewPasswordGenerator creates a new password generator
//
// The password generator creates a string using alphanumeric characters and symbols.
func NewPasswordGenerator() Generator {
	return newStringGenerator(lowercase + uppercase + numbers + symbols)
}
