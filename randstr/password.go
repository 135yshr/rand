package randstr

import (
	"math/rand"
	"strings"
	"time"
)

// PasswordPolicy is a function that checks if a password is valid
type PasswordPolicy func(string) bool

var (
	// PasswordPolicyLength is a password policy that checks if the password is a certain length
	PasswordPolicyLength = func(length int) PasswordPolicy {
		return func(password string) bool {
			return len(password) >= length
		}
	}

	// PasswordPolicyLowercase is a password policy that checks if the password contains a lowercase character
	PasswordPolicyLowercase = func() PasswordPolicy {
		return func(password string) bool {
			return strings.ContainsAny(password, lowercase)
		}
	}

	// PasswordPolicyUppercase is a password policy that checks if the password contains an uppercase character
	PasswordPolicyUppercase = func() PasswordPolicy {
		return func(password string) bool {
			return strings.ContainsAny(password, uppercase)
		}
	}

	// PasswordPolicyNumbers is a password policy that checks if the password contains a number
	PasswordPolicyNumbers = func() PasswordPolicy {
		return func(password string) bool {
			return strings.ContainsAny(password, numbers)
		}
	}

	// PasswordPolicySymbols is a password policy that checks if the password contains a symbol
	PasswordPolicySymbols = func() PasswordPolicy {
		return func(password string) bool {
			return strings.ContainsAny(password, symbols)
		}
	}
)

type passwordGenerator struct {
	stringGenerator
	policies []PasswordPolicy
}

// NewPasswordGenerator creates a new password generator
//
// The password generator creates a string using alphanumeric characters and symbols.
func NewPasswordGenerator(policies ...PasswordPolicy) Generator {
	return &passwordGenerator{
		stringGenerator: stringGenerator{
			letters: []rune(lowercase + uppercase + numbers + symbols),
			r:       rand.NewSource(time.Now().UnixNano()),
		},
		policies: policies,
	}
}

// Generate generates a password
func (g *passwordGenerator) Generate(length int) string {
	password := g.stringGenerator.Generate(length)
	for _, policy := range g.policies {
		if !policy(password) {
			return g.Generate(length)
		}
	}
	return password
}
