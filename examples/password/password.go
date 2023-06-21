package main

import (
	"fmt"

	"github.com/135yshr/rand/randstr"
)

func main() {
	gen := randstr.NewPasswordGenerator()
	fmt.Printf("Generated password 1: %s\n", gen.Generate(12))
	fmt.Printf("Generated password 2: %s\n", gen.Generate(12))
	fmt.Printf("Generated password 3: %s\n", gen.Generate(12))
}
