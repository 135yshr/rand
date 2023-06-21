package randstr

type passwordGenerator StringGenerator

// NewPasswordGenerator creates a new password generator
func NewPasswordGenerator() Generator {
	return (*passwordGenerator)(NewStringGenerator(lowercase + uppercase + numbers + symbols))
}

// Generate generates a password
func (g *passwordGenerator) Generate(n int) string {
	return generate(StringGenerator(*g), n)
}
