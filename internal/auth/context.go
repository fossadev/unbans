package auth

import (
	"context"

	"github.com/fossadev/unbans/internal/features/token"
)

type contextKey string

const (
	jwtClaimsKey = contextKey("auth.jwtClaims")
	jwtTokenKey  = contextKey("auth.jwtToken")
)

func GetJWTClaims(ctx context.Context) *token.JWTClaims {
	if claims, ok := ctx.Value(jwtClaimsKey).(*token.JWTClaims); ok {
		return claims
	}
	return nil
}

func HasJWTClaims(ctx context.Context) bool {
	return GetJWTClaims(ctx) != nil
}

func WithJWTClaims(ctx context.Context, claims *token.JWTClaims) context.Context {
	return context.WithValue(ctx, jwtClaimsKey, claims)
}

func GetJWTToken(ctx context.Context) string {
	if token, ok := ctx.Value(jwtTokenKey).(string); ok {
		return token
	}
	return ""
}

func HasJWTToken(ctx context.Context) bool {
	return GetJWTToken(ctx) != ""
}

func WithJWTToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, jwtTokenKey, token)
}
