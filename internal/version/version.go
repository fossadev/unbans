package version

import (
	"github.com/fossadev/unbans/internal/cache"
	"github.com/fossadev/unbans/internal/controller"
	"github.com/fossadev/unbans/internal/db"
	"github.com/fossadev/unbans/internal/features"
	"github.com/fossadev/unbans/internal/logger"
	"github.com/go-chi/chi"
)

type Version struct {
	*chi.Mux
	controllerCfg *controller.Config
}

type Config struct {
	Cache    cache.Cache
	DB       *db.DB
	Features *features.Features
	Log      logger.Logger
}

func New(conf *Config) *Version {
	r := chi.NewRouter()

	return &Version{
		Mux: r,
		controllerCfg: &controller.Config{
			Cache:    conf.Cache,
			DB:       conf.DB,
			Features: conf.Features,
			Log:      conf.Log,
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
