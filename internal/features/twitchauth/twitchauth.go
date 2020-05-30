package twitchauth

import (
	"context"
	"time"

	"github.com/fossadev/unbans/internal/cache"
	"github.com/fossadev/unbans/internal/config"
	"github.com/fossadev/unbans/internal/db"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/twitch"
)

const (
	stateTokenExpiry = time.Minute * 30
)

var StateAttemptMaxErr = errors.New("state attempts maxed")

type TwitchAuthFeature interface {
	ExchangeCode(ctx context.Context, code, state string) (token *oauth2.Token, next string, err error)
	GetAuthURL(next string) (redirectURL string, err error)
}

type stateData struct {
	NextURL string `json:"next_url"`
}

type twitchAuthFeature struct {
	cache cache.Cache
	db    *db.DB
	oauth *oauth2.Config
}

func New(cfg *config.TwitchConfig, dbImpl *db.DB, cacheImpl cache.Cache) TwitchAuthFeature {
	return &twitchAuthFeature{
		cache: cacheImpl,
		db:    dbImpl,
		oauth: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  cfg.RedirectURL,
			Scopes:       config.TwitchScopes,
			Endpoint:     twitch.Endpoint,
		},
	}
}
