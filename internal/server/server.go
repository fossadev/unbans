package server

import (
	"context"
	"net/http"

	"github.com/fossadev/unbans/internal/config"
	"github.com/fossadev/unbans/internal/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

type Server interface {
	Start()
	Stop(ctx context.Context) error
}

type server struct {
	log      logger.Logger
	srv      *http.Server
	quitChan chan bool
}

func New(cfg *config.Config, log logger.Logger) Server {
	s := &server{
		log:      log,
		quitChan: make(chan bool, 1),
	}

	s.srv = &http.Server{
		Addr:    cfg.WebHost,
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

	r.Use(middleware.Recoverer)

	return r
}
