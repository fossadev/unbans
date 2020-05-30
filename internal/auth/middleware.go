package auth

import (
	"net/http"
	"strings"

	"github.com/fossadev/unbans/internal/encoder"
	"github.com/fossadev/unbans/internal/features"
	"github.com/fossadev/unbans/internal/logger"
)

func Middleware(ft *features.Features, log logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			header := req.Header.Get("authorization")
			if header == "" {
				next.ServeHTTP(w, req)
				return
			}

			req.Header.Del("authorization")

			rw := encoder.NewResponseWriter(w, req, log)

			spl := strings.SplitN(header, " ", 2)
			if len(spl) != 2 || spl[0] == "Bearer" {
				rw.BadRequest("Authorization header must be in format: Bearer <token>")
				return
			}

			token := strings.TrimSpace(spl[1])
			if token == "" {
				next.ServeHTTP(w, req)
				return
			}

			ctx := req.Context()
			claims, err := ft.Token.ValidateJWT(token)
			if err == nil && claims != nil {
				ctx = WithJWTToken(ctx, token)
				ctx = WithJWTClaims(ctx, claims)
			}

			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}
