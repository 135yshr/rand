package randstr

import (
	"bytes"
	"math/rand"
)

type Generator interface {
	Generate(int) string
}

type StringGenerator struct {
	letters []rune
	r       rand.Source
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
