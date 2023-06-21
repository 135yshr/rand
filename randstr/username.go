package randstr

import (
	"bytes"
	"math/rand"
	"time"
)

type userNameGenerator struct {
	letters []rune
	r       rand.Source
}

// NewUserNameGenerator returns a new UserName generator.
func NewUserNameGenerator() Generator {
	return &userNameGenerator{
		letters: []rune(lowercase + uppercase + numbers),
		r:       rand.NewSource(time.Now().UnixNano()),
	}
}

// Generate generates a random user name.
func (g *userNameGenerator) Generate(n int) string {
	var bb bytes.Buffer
	bb.Grow(n)

	l := len(g.letters)
	for i := 0; i < n; i++ {
		bb.WriteRune(g.letters[rand.Intn(l)])
	}
	return bb.String()
}
