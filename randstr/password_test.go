package randstr_test

import (
	"testing"

	"github.com/135yshr/rand/randstr"
	"github.com/stretchr/testify/require"
)

// TestPasswordGeneratorGenerate tests the password generator
func TestPasswordGeneratorGenerate(t *testing.T) {
	type args struct {
		length   int
		policies []randstr.PasswordPolicy
	}
	type want struct {
		length  int
		pattern string
	}
	tests := map[string]struct {
		args args
		want want
	}{
		"should return a 5 characters string of mixed case": {
			args: args{
				length: 5,
			},
			want: want{
				length:  5,
				pattern: "^[ -~]{5}",
			},
		},
		"should return a 10 characters string of mixed case": {
			args: args{
				length: 10,
			},
			want: want{
				length:  10,
				pattern: "^[ -~]{10}",
			},
		},
		"should return a 5 characters string that definitely contains uppercase letters": {
			args: args{
				length: 5,
				policies: []randstr.PasswordPolicy{
					randstr.PasswordPolicyUppercase(),
				},
			},
			want: want{
				length:  5,
				pattern: "[A-Z]",
			},
		},
		"should return a 5 characters string that definitely contains lowercase letters": {
			args: args{
				length: 5,
				policies: []randstr.PasswordPolicy{
					randstr.PasswordPolicyLowercase(),
				},
			},
			want: want{
				length:  5,
				pattern: "[a-z]",
			},
		},
		"should return a 5 characters string that definitely contains numbers": {
			args: args{
				length: 5,
				policies: []randstr.PasswordPolicy{
					randstr.PasswordPolicyNumbers(),
				},
			},
			want: want{
				length:  5,
				pattern: "[0-9]",
			},
		},
		"should return a 5 characters string that definitely contains symbols": {
			args: args{
				length: 5,
				policies: []randstr.PasswordPolicy{
					randstr.PasswordPolicySymbols(),
				},
			},
			want: want{
				length:  5,
				pattern: "[!-/:-@[-`{-~]",
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sut := randstr.NewPasswordGenerator(tt.args.policies...)
			require.NotNil(t, sut)

			str := sut.Generate(tt.args.length)
			require.Len(t, str, tt.want.length)
			require.Regexp(t, tt.want.pattern, str)
			t.Logf("return: %s", str)
		})
	}
}
