package helpers

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/pkg/errors"
)

func GenerateRandomToken(length int) (string, error) {
	bs := make([]byte, int(float64(length/2)))
	if _, err := rand.Read(bs); err != nil {
		return "", errors.Wrap(err, "failed to generate random token")
	}
	return hex.EncodeToString(bs)[0 : length-1], nil
}
