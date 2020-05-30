package models

import (
	"context"
	"time"

	"github.com/go-pg/pg/v9"
)

type Channel struct {
	tableName struct{} `pg:"channels"`

	ID              int       `pg:"id"`
	Login           string    `pg:"login"`
	DisplayName     string    `pg:"display_name"`
	Avatar          string    `pg:"avatar"`
	Provider        string    `pg:"provider"`
	ProviderID      string    `pg:"provider_id"`
	BroadcasterType string    `pg:"broadcaster_type"`
	Slug            string    `pg:"slug"`
	AccessToken     string    `pg:"access_token"`
	RefreshToken    string    `pg:"refresh_token"`
	TokenExpires    time.Time `pg:"token_expires"`
	CreatedAt       time.Time `pg:"created_at"`
	UpdatedAt       time.Time `pg:"updated_at"`
}

var _ pg.BeforeInsertHook = (*Channel)(nil)

func (m *Channel) BeforeInsert(ctx context.Context) (context.Context, error) {
	now := time.Now().UTC()
	m.CreatedAt = now
	m.UpdatedAt = now
	return ctx, nil
}

var _ pg.BeforeUpdateHook = (*Channel)(nil)

func (m *Channel) BeforeUpdate(ctx context.Context) (context.Context, error) {
	m.UpdatedAt = time.Now().UTC()
	return ctx, nil
}
