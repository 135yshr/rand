package randstr_test

import (
	"testing"

	"github.com/135yshr/rand/randstr"
	"github.com/stretchr/testify/require"
)

// TestCustomGeneratorGenerate tests creating a random string.
func TestCustomGeneratorGenerate(t *testing.T) {
	type args struct {
		letters string
		length  int
	}
	type want struct {
		length  int
		pattern string
	}
	tests := map[string]struct {
		args args
		want want
	}{
		"should return a 5 character string using 'abcdefg'": {
			args: args{
				letters: "abcdefg",
				length:  5,
			},
			want: want{
				length:  5,
				pattern: "[a-g]{5}",
			},
		},
		"should return a 5 character string using 'abcdefg012345'": {
			args: args{
				letters: "abcdefg012345",
				length:  5,
			},
			want: want{
				length:  5,
				pattern: "[a-g0-5]{5}",
			},
		},
		"should return a 10 character string using 'abcdefg012345'": {
			args: args{
				letters: "abcdefg012345",
				length:  10,
			},
			want: want{
				length:  10,
				pattern: "[a-g0-5]{10}",
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sut := randstr.NewCustomGenerator(tt.args.letters)
			require.NotNil(t, sut)

			str := sut.Generate(tt.args.length)
			require.Len(t, str, tt.want.length)
			require.Regexp(t, tt.want.pattern, str)
			t.Logf("return: %s", str)
		})
	}
}
