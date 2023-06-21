package randstr

import (
	"math/rand"
	"time"
)

type userNameGenerator StringGenerator

// NewUserNameGenerator returns a new UserName generator.
func NewUserNameGenerator() Generator {
	return &userNameGenerator{
		letters: []rune(lowercase + uppercase + numbers),
		r:       rand.NewSource(time.Now().UnixNano()),
	}
}

// Generate generates a random user name.
func (g *userNameGenerator) Generate(n int) string {
	return generate(StringGenerator(*g), n)
}
