# randstr

[![go](https://github.com/135yshr/rand/actions/workflows/go.yml/badge.svg)](https://github.com/135yshr/rand/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/135yshr/rand)](https://goreportcard.com/report/github.com/135yshr/rand)
[![GoDoc](http://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://godoc.org/github.com/135yshr/rand)
[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

This is a library for generating various random data to use in the program.

## Prerequisites

- [Go](https://go.dev/): 1.18+

## Installation

```bash
go get -u github.com/135yshr/rand
```

## Lean more

### How to generate random ID

The ID generator creates a string using alphanumeric characters.

```go
import (
	"fmt"

	"github.com/135yshr/rand/randstr"
)

func main() {
	gen := randstr.NewUserNameGenerator()
	fmt.Printf("Generated ID: %s\n", gen.Generate(8))
}
```

### How to generate random Password

```go
import (
	"fmt"

	"github.com/135yshr/rand/randstr"
)

func main() {
	gen := randstr.NewPasswordGenerator()
	fmt.Printf("Generated password: %s\n", gen.Generate(12))
}
```

### How to generate random string

When you want to use specified characters, please use the custom generator.

```go
import (
	"fmt"

	"github.com/135yshr/rand/randstr"
)

func main() {
	gen := randstr.NewCustomGenerator("abcdef01234567890-/")
	fmt.Printf("Generated random string: %s\n", gen.Generate(10))
}
```

## Contributing

This project is an open-source endeavor that thrives on your active participation. We're always on the lookout for individuals interested in contributing to the project's growth. If you have any ideas or improvements, no matter how small, they are welcome. Feel free to submit a [pull request](https://github.com/135yshr/rand/pulls) at any time. We're eagerly awaiting your collaboration!

## License

This project is released under the MIT license. See the [LICENSE](LICENSE) file for details.
