package randstr_test

import (
	"testing"

	"github.com/135yshr/rand/randstr"
	"github.com/stretchr/testify/require"
)

// TestGeneratel tests the Dummy function.
func TestGenerate(t *testing.T) {
	require.Equal(t, "dummy", randstr.Generate())
}
