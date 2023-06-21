package randstr

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	symbols   = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

type defaultGenerator StringGenerator

// NewDefaultGenerator returns a new Default generator.
func NewDefaultGenerator() Generator {
	return (*defaultGenerator)(NewStringGenerator(lowercase + uppercase + numbers + symbols))
}

// Generate generates a random string.
func (g *defaultGenerator) Generate(n int) string {
	return generate(StringGenerator(*g), n)
}
