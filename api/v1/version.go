package v1

import (
	"github.com/fossadev/unbans/api/v1/user"
	"github.com/fossadev/unbans/internal/version"
)

func New(conf *version.Config) *version.Version {
	v := version.New(conf)

	v.Register(user.New)

	return v
}
