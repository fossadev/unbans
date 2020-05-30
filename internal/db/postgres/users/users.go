package users

import (
	"context"

	"github.com/fossadev/unbans/internal/db"
	"github.com/fossadev/unbans/internal/models"
	"github.com/go-pg/pg/v9"
)

type users struct {
	client *pg.DB
}

func New(client *pg.DB) db.Users {
	return &users{client: client}
}

func (m *users) UpsertUser(ctx context.Context, model *models.User) error {
	_, err := m.client.WithContext(ctx).Model(model).
		OnConflict("(provider, provider_id) DO UPDATE").
		Set("channel_id = EXCLUDED.channel_id").
		Set("login = EXCLUDED.login").
		Set("display_name = EXCLUDED.display_name").
		Set("avatar = EXCLUDED.avatar").
		Set("provider = EXCLUDED.provider").
		Set("provider_id = EXCLUDED.provider_id").
		Set("type = EXCLUDED.type").
		Set("last_login = EXCLUDED.last_login").
		Set("updated_at = EXCLUDED.updated_at").
		Returning("*").
		Insert()
	return err
}
