package randstr

import (
	"bytes"
	"math/rand"
	"time"
)

type Generator interface {
	Generate(int) string
}

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
)

type defaultGenerator struct {
	letters []rune
	r       rand.Source
}

// NewDefaultGenerator returns a new Default generator.
func NewDefaultGenerator() Generator {
	seed := time.Now().UnixNano()
	return &defaultGenerator{
		letters: []rune(lowercase),
		r:       rand.NewSource(seed),
	}
}

// Generate generates a random string.
func (g *defaultGenerator) Generate(n int) string {
	var bb bytes.Buffer
	bb.Grow(n)

	l := len(g.letters)
	for i := 0; i < n; i++ {
		bb.WriteRune(g.letters[rand.Intn(l)])
	}
	return bb.String()
}
