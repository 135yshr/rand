package randstr

import (
	"bytes"
	"math/rand"
	"time"
)

type Generator interface {
	Generate(int) string
}

type StringGenerator struct {
	r       rand.Source
	letters []rune
}

// NewStringGenerator creates a new string generator
func NewStringGenerator(letters string) *StringGenerator {
	return &StringGenerator{
		letters: []rune(letters),
		r:       rand.NewSource(time.Now().UnixNano()),
	}
}

func generate(s StringGenerator, n int) string {
	var bb bytes.Buffer
	bb.Grow(n)

	l := len(s.letters)
	for i := 0; i < n; i++ {
		bb.WriteRune(s.letters[rand.Intn(l)])
	}
	return bb.String()
}
