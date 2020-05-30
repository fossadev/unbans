package auth

import (
	"github.com/fossadev/unbans/api/auth/twitch"
	"github.com/fossadev/unbans/internal/version"
)

func New(conf *version.Config) *version.Version {
	v := version.New(conf)

	v.Register(twitch.New)

	return v
}
