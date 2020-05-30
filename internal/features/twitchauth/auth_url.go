package twitchauth

import (
	"encoding/json"

	"github.com/fossadev/unbans/internal/helpers"
	"github.com/pkg/errors"
)

func (f *twitchAuthFeature) GetAuthURL(next string) (string, error) {
	data := &stateData{
		NextURL: next,
	}

	bs, err := json.Marshal(data)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal state data")
	}

	for i := 0; i < 25; i++ {
		stateToken, err := helpers.GenerateRandomToken(50)
		if err != nil {
			return "", errors.Wrap(err, "failed to generate random token")
		}

		exists, err := f.cache.AttemptStateToken(stateToken, string(bs), stateTokenExpiry)
		if err != nil {
			return "", err
		}

		if !exists {
			return f.oauth.AuthCodeURL(stateToken), nil
		}
	}

	return "", StateAttemptMaxErr
}
