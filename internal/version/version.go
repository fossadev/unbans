package version

import (
	"github.com/fossadev/unbans/internal/auth"
	"github.com/fossadev/unbans/internal/cache"
	"github.com/fossadev/unbans/internal/config"
	"github.com/fossadev/unbans/internal/controller"
	"github.com/fossadev/unbans/internal/db"
	"github.com/fossadev/unbans/internal/features"
	"github.com/fossadev/unbans/internal/logger"
	"github.com/fossadev/unbans/internal/twitchapi"
	"github.com/go-chi/chi"
)

type Version struct {
	*chi.Mux
	controllerCfg *controller.Config
}

type Config struct {
	Cache     cache.Cache
	Config    *config.Config
	DB        *db.DB
	Features  *features.Features
	Log       logger.Logger
	TwitchAPI twitchapi.TwitchAPI
}

func New(conf *Config) *Version {
	r := chi.NewRouter()

	r.Use(auth.Middleware(conf.Features, conf.Log))

	return &Version{
		Mux: r,
		controllerCfg: &controller.Config{
			Cache:     conf.Cache,
			Config:    conf.Config,
			DB:        conf.DB,
			Features:  conf.Features,
			Log:       conf.Log,
			TwitchAPI: conf.TwitchAPI,
		},
	}
}

func (v *Version) Register(controllerFactory func(*controller.Controller)) {
	c := controller.New(v.controllerCfg)
	controllerFactory(c)

	for _, r := range c.Routes() {
		r.Inject(v.Mux)
	}
}
