package controller

import (
	"context"
	"net/http"

	"github.com/fossadev/unbans/internal/encoder"
	"github.com/fossadev/unbans/internal/logger"
	"github.com/go-chi/chi"
)

type handlerFunc func(ctx context.Context, w encoder.ResponseWriter, req *http.Request)

type Route struct {
	method     string
	pattern    string
	handler    handlerFunc
	controller *Controller
	log        logger.Logger
}

func (r *Route) Inject(mux *chi.Mux) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		rw := encoder.NewResponseWriter(w, req, r.log)
		ctx := req.Context()

		r.handler(ctx, rw, req)
	})

	mux.Group(func(g chi.Router) {
		switch r.method {
		case http.MethodGet:
			g.Get(r.pattern, handler)
		case http.MethodPost:
			g.Post(r.pattern, handler)
		case http.MethodPatch:
			g.Patch(r.pattern, handler)
		case http.MethodPut:
			g.Put(r.pattern, handler)
		case http.MethodDelete:
			g.Delete(r.pattern, handler)
		}
	})
}
