package controller

import (
	"net/http"

	"github.com/fossadev/unbans/internal/cache"
	"github.com/fossadev/unbans/internal/db"
	"github.com/fossadev/unbans/internal/features"
	"github.com/fossadev/unbans/internal/logger"
)

type Controller struct {
	routes []*Route

	Cache    cache.Cache
	DB       *db.DB
	Features *features.Features
	Log      logger.Logger
}

type Config struct {
	Cache    cache.Cache
	DB       *db.DB
	Features *features.Features
	Log      logger.Logger
}

func New(conf *Config) *Controller {
	return &Controller{
		Cache:    conf.Cache,
		DB:       conf.DB,
		Features: conf.Features,
		Log:      conf.Log,
	}
}

func (c *Controller) Routes() []*Route {
	return c.routes
}

func (c *Controller) Get(pattern string, handler handlerFunc) *Route {
	return c.registerRoute(http.MethodGet, pattern, handler)
}

func (c *Controller) Post(pattern string, handler handlerFunc) *Route {
	return c.registerRoute(http.MethodPost, pattern, handler)
}

func (c *Controller) Put(pattern string, handler handlerFunc) *Route {
	return c.registerRoute(http.MethodPut, pattern, handler)
}

func (c *Controller) Patch(pattern string, handler handlerFunc) *Route {
	return c.registerRoute(http.MethodPatch, pattern, handler)
}

func (c *Controller) Delete(pattern string, handler handlerFunc) *Route {
	return c.registerRoute(http.MethodDelete, pattern, handler)
}

func (c *Controller) registerRoute(method, pattern string, handler handlerFunc) *Route {
	r := &Route{
		method:     method,
		pattern:    pattern,
		handler:    handler,
		controller: c,
		log:        c.Log,
	}
	c.routes = append(c.routes, r)
	return r
}
