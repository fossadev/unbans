package models

import (
	"context"
	"time"

	"github.com/go-pg/pg/v9"
)

type User struct {
	tableName struct{} `pg:"users"`

	ID          int       `pg:"id"`
	ChannelID   int       `pg:"channel_id"`
	Userlevel   int       `pg:"userlevel"`
	Login       string    `pg:"login"`
	DisplayName string    `pg:"display_name"`
	Avatar      string    `pg:"avatar"`
	Provider    string    `pg:"provider"`
	ProviderID  string    `pg:"provider_id"`
	Type        string    `pg:"type"`
	LastLogin   time.Time `pg:"last_login"`
	CreatedAt   time.Time `pg:"created_at"`
	UpdatedAt   time.Time `pg:"updated_at"`
}

var _ pg.BeforeInsertHook = (*User)(nil)

func (m *User) BeforeInsert(ctx context.Context) (context.Context, error) {
	now := time.Now().UTC()
	m.CreatedAt = now
	m.UpdatedAt = now
	if m.LastLogin.IsZero() {
		m.LastLogin = now
	}
	return ctx, nil
}

var _ pg.BeforeUpdateHook = (*User)(nil)

func (m *User) BeforeUpdate(ctx context.Context) (context.Context, error) {
	m.UpdatedAt = time.Now().UTC()
	return ctx, nil
}
