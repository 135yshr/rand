package randstr_test

import (
	"testing"

	"github.com/135yshr/rand/randstr"
	"github.com/stretchr/testify/require"
)

// TestIDGeneratorGenerate tests creating a random user name.
func TestIDGeneratorGenerate(t *testing.T) {
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
				pattern: "[a-zA-Z0-9]{5}",
			},
		},
		"should return a string of 10 characters": {
			args: args{
				length: 10,
			},
			want: want{
				length:  10,
				pattern: "[a-zA-Z0-9]{10}",
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sut := randstr.NewIDGenerator()
			require.NotNil(t, sut)

			str := sut.Generate(tt.args.length)
			require.Len(t, str, tt.want.length)
			require.Regexp(t, tt.want.pattern, str)
			t.Logf("return: %s", str)
		})
	}
}
