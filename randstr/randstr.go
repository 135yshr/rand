package randstr

type Generator interface {
	Generate(int) string
}

type defaultGenerator struct{}

// NewDefaultGenerator returns a new Default generator.
func NewDefaultGenerator() Generator {
	return &defaultGenerator{}
}

// Generate generates a random string.
func (*defaultGenerator) Generate(int) string {
	return "dummy"
}
