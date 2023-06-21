package randstr

type passwordGenerator stringGenerator

// NewPasswordGenerator creates a new password generator
func NewPasswordGenerator() Generator {
	return newStringGenerator(lowercase + uppercase + numbers + symbols)
}
