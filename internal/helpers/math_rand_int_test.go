package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMathRandInt(t *testing.T) {
	generated := make(map[int]bool)
	for i := 0; i < 10; i++ {
		x := MathRandInt(1, 9999)
		if _, ok := generated[x]; ok {
			t.Errorf("failed to ensure random numbers after %d attempts, got %d more than once!", i+1, x)
			break
		}
		assert.Equal(t, true, x >= 1, "number must be within lower bound")
		assert.Equal(t, true, x <= 9999, "number must be within upper bound")
		generated[x] = true
	}
}
