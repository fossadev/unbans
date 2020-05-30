package twitchauth

import (
	"context"

	"github.com/fossadev/unbans/internal/cache"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

var StateNotFoundErr = errors.New("state not found")

func (f *twitchAuthFeature) ExchangeCode(ctx context.Context, code, state string) (*oauth2.Token, string, error) {
	var payload stateData
	if err := f.cache.GetStatePayload(state, &payload); err != nil {
		if err == cache.ErrKeyNotFound {
			return nil, "", StateNotFoundErr
		}
		return nil, "", err
	}

	token, err := f.oauth.Exchange(ctx, code)
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to exchange code")
	}

	return token, "", nil
}
