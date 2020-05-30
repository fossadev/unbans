package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fossadev/unbans/internal/cache"
	"github.com/fossadev/unbans/internal/config"
	"github.com/fossadev/unbans/internal/db"
	"github.com/pkg/errors"
)

const jwtTokenLifetime = time.Hour * 12

type TokenFeature interface {
	CreateJWT(userID, channelID int, provider, providerID string) (token string, err error)
	ValidateJWT(raw string) (parsed *JWTClaims, err error)
}

type tokenFeature struct {
	cache     cache.Cache
	db        *db.DB
	jwtSecret []byte
}

type JWTClaims struct {
	jwt.StandardClaims
	ChannelID  int    `json:"channel_id"`
	Provider   string `json:"provider"`
	ProviderID string `json:"provider_id"`
	UserID     int    `json:"user_id"`
}

func New(cacheImpl cache.Cache, cfg *config.Config, dbImpl *db.DB) TokenFeature {
	return &tokenFeature{
		cache:     cacheImpl,
		db:        dbImpl,
		jwtSecret: []byte(cfg.JWTSecret),
	}
}

func (f *tokenFeature) ValidateJWT(raw string) (*JWTClaims, error) {
	claims := &JWTClaims{}
	t, err := jwt.ParseWithClaims(raw, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return f.jwtSecret, nil
	})
	if err != nil || !t.Valid {
		return nil, errors.Wrap(err, "invalid token")
	}

	return claims, nil
}

func (f *tokenFeature) CreateJWT(userID, channelID int, provider, providerID string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTClaims{
		UserID:     userID,
		Provider:   provider,
		ProviderID: providerID,
		ChannelID:  channelID,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "fossadev_unbans",
			ExpiresAt: time.Now().Add(jwtTokenLifetime).Unix(),
		},
	})

	token, err := t.SignedString(f.jwtSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}
