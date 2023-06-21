package randstr

import (
	"math/rand"
	"time"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	symbols   = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

type defaultGenerator StringGenerator

// NewDefaultGenerator returns a new Default generator.
func NewDefaultGenerator() Generator {
	return &defaultGenerator{
		letters: []rune(lowercase + uppercase + numbers + symbols),
		r:       rand.NewSource(time.Now().UnixNano()),
	}
}

// Generate generates a random string.
func (g *defaultGenerator) Generate(n int) string {
	return generate(StringGenerator(*g), n)
}
