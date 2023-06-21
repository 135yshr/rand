package randstr_test

import (
	"testing"

	"github.com/135yshr/rand/randstr"
	"github.com/stretchr/testify/require"
)

// TestGenerate tests the Dummy function.
func TestGenerate(t *testing.T) {
	g := randstr.NewDefaultGenerator()
	require.NotNil(t, g)

	str := g.Generate(5)
	require.Len(t, str, 5)
	require.Regexp(t, "[a-z]+", str)
}
