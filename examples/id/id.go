package main

import (
	"fmt"

	"github.com/135yshr/rand/randstr"
)

func main() {
	gen := randstr.NewIDGenerator()
	fmt.Printf("Generated username 1: %s\n", gen.Generate(8))
	fmt.Printf("Generated username 2: %s\n", gen.Generate(8))
	fmt.Printf("Generated username 3: %s\n", gen.Generate(8))
}
