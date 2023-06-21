package randstr_test

import (
	"testing"

	"github.com/135yshr/rand/randstr"
	"github.com/stretchr/testify/require"
)

// TestDefaultGeneratorGenerate tests creating a random string.
func TestDefaultGeneratorGenerate(t *testing.T) {
	type args struct {
		length int
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
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sut := randstr.NewDefaultGenerator()
			require.NotNil(t, sut)

			str := sut.Generate(tt.args.length)
			require.Len(t, str, tt.want.length)
			require.Regexp(t, tt.want.pattern, str)
			t.Logf("return: %s", str)
		})
	}
}

// TestUserNameGeneratorGenerate tests creating a random user name.
func TestUserNameGeneratorGenerate(t *testing.T) {
	type args struct {
		length int
	}
	type want struct {
		length  int
		pattern string
	}
	tests := map[string]struct {
		args args
		want want
	}{
		"should return a string of 5 characters": {
			args: args{
				length: 5,
			},
			want: want{
				length:  5,
				pattern: "[a-zA-Z0-0]{5}",
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sut := randstr.NewUserNameGenerator()
			require.NotNil(t, sut)

			str := sut.Generate(tt.args.length)
			require.Len(t, str, tt.want.length)
			require.Regexp(t, tt.want.pattern, str)
			t.Logf("return: %s", str)
		})
	}
}
