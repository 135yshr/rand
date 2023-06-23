package randstr

import (
	"bytes"
	"math/rand"
	"time"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	symbols   = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

type Generator interface {
	Generate(int) string
}

type stringGenerator struct {
	r       rand.Source
	letters []rune
}

func newStringGenerator(letters string) Generator {
	return &stringGenerator{
		letters: []rune(letters),
		r:       rand.NewSource(time.Now().UnixNano()),
	}
}

// NewDefaultGenerator returns a new Custom generator.
func NewCustomGenerator(letters string) Generator {
	return newStringGenerator(letters)
}

// NewDefaultGenerator returns a new generator
func (s *stringGenerator) Generate(n int) string {
	var bb bytes.Buffer
	bb.Grow(n)

	l := len(s.letters)
	for i := 0; i < n; i++ {
		bb.WriteRune(s.letters[rand.Intn(l)])
	}
	return bb.String()
}
