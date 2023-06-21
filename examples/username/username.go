package main

import (
	"fmt"

	"github.com/135yshr/rand/randstr"
)

func main() {
	gen := randstr.NewUserNameGenerator()
	fmt.Printf("Generate username 1: %s\n", gen.Generate(8))
	fmt.Printf("Generate username 2: %s\n", gen.Generate(8))
	fmt.Printf("Generate username 3: %s\n", gen.Generate(8))
}
