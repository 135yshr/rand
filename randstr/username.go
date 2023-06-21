package randstr

type userNameGenerator StringGenerator

// NewUserNameGenerator returns a new UserName generator.
func NewUserNameGenerator() Generator {
	return (*userNameGenerator)(NewStringGenerator(lowercase + uppercase + numbers))
}

// Generate generates a random user name.
func (g *userNameGenerator) Generate(n int) string {
	return generate(StringGenerator(*g), n)
}
