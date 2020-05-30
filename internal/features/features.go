package features

import (
	"github.com/fossadev/unbans/internal/cache"
	"github.com/fossadev/unbans/internal/config"
	"github.com/fossadev/unbans/internal/db"
	"github.com/fossadev/unbans/internal/features/channel"
	"github.com/fossadev/unbans/internal/features/token"
	"github.com/fossadev/unbans/internal/features/twitchauth"
	"github.com/fossadev/unbans/internal/features/user"
	"github.com/fossadev/unbans/internal/logger"
)

type Features struct {
	Channel    channel.ChannelFeature
	Token      token.TokenFeature
	TwitchAuth twitchauth.TwitchAuthFeature
	User       user.UserFeature
}

func New(cacheImpl cache.Cache, cfg *config.Config, dbImpl *db.DB, log logger.Logger) *Features {
	return &Features{
		Channel:    channel.New(dbImpl, log),
		Token:      token.New(cacheImpl, cfg, dbImpl),
		TwitchAuth: twitchauth.New(cfg.Twitch, dbImpl, cacheImpl),
		User:       user.New(dbImpl),
	}
}
