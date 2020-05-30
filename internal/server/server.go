package server

import (
	"context"
	"net/http"

	"github.com/fossadev/unbans/api/auth"
	v1 "github.com/fossadev/unbans/api/v1"
	"github.com/fossadev/unbans/internal/cache"
	"github.com/fossadev/unbans/internal/config"
	"github.com/fossadev/unbans/internal/db"
	"github.com/fossadev/unbans/internal/features"
	"github.com/fossadev/unbans/internal/logger"
	"github.com/fossadev/unbans/internal/twitchapi"
	"github.com/fossadev/unbans/internal/version"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

type Server interface {
	Start()
	Stop(ctx context.Context) error
}

type Config struct {
	Cache     cache.Cache
	Config    *config.Config
	DB        *db.DB
	Features  *features.Features
	Log       logger.Logger
	TwitchAPI twitchapi.TwitchAPI
}

type server struct {
	conf     *Config
	log      logger.Logger
	srv      *http.Server
	quitChan chan bool
}

func New(conf *Config) Server {
	s := &server{
		conf:     conf,
		log:      conf.Log,
		quitChan: make(chan bool, 1),
	}

	s.srv = &http.Server{
		Addr:    conf.Config.WebHost,
		Handler: s.handler(),
	}

	return s
}

func (s *server) Start() {
	go func() {
		s.log.Info("server started", zap.String("addr", s.srv.Addr))
		if err := s.srv.ListenAndServe(); err != nil {
			select {
			case <-s.quitChan:
			default:
				s.log.Error("server closed", err)
			}
		}
	}()
}

func (s *server) Stop(ctx context.Context) error {
	s.quitChan <- true
	return s.srv.Shutdown(ctx)
}

func (s *server) handler() *chi.Mux {
	r := chi.NewRouter()

	versionConf := &version.Config{
		Cache:     s.conf.Cache,
		Config:    s.conf.Config,
		DB:        s.conf.DB,
		Features:  s.conf.Features,
		Log:       s.conf.Log,
		TwitchAPI: s.conf.TwitchAPI,
	}

	r.Use(middleware.Recoverer)

	r.Mount("/auth", auth.New(versionConf))
	r.Mount("/v1", v1.New(versionConf))

	return r
}
