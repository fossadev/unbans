package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomToken(t *testing.T) {
	pastTokens := make(map[string]bool)
	for i := 0; i < 100; i++ {
		token, err := GenerateRandomToken(24)
		assert.Nil(t, err)

		_, ok := pastTokens[token]
		assert.Equal(t, false, ok, "tokens are unique")
		pastTokens[token] = true
	}
}
