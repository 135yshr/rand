package main

import (
	"fmt"

	"github.com/135yshr/rand/randstr"
)

func main() {
	gen := randstr.NewCustomGenerator("abcdef01234567890-/")
	fmt.Printf("Generate random string: %s\n", gen.Generate(10))
	fmt.Printf("Generate random string: %s\n", gen.Generate(10))
	fmt.Printf("Generate random string: %s\n", gen.Generate(10))
}
