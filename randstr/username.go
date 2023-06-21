package randstr

// NewUserNameGenerator returns a new UserName generator.
func NewUserNameGenerator() Generator {
	return newStringGenerator(lowercase + uppercase + numbers)
}
