package channels

import (
	"context"

	"github.com/fossadev/unbans/internal/db"
	"github.com/fossadev/unbans/internal/models"
	"github.com/go-pg/pg/v9"
)

type channels struct {
	client *pg.DB
}

func New(client *pg.DB) db.Channels {
	return &channels{client: client}
}

func (m *channels) UpsertChannel(ctx context.Context, model *models.Channel) error {
	_, err := m.client.WithContext(ctx).Model(model).
		OnConflict("(provider, provider_id) DO UPDATE").
		Set("login = EXCLUDED.login").
		Set("display_name = EXCLUDED.display_name").
		Set("avatar = EXCLUDED.avatar").
		Set("provider = EXCLUDED.provider").
		Set("provider_id = EXCLUDED.provider_id").
		Set("broadcaster_type = EXCLUDED.broadcaster_type").
		Set("slug = EXCLUDED.slug").
		Set("access_token = EXCLUDED.access_token").
		Set("refresh_token = EXCLUDED.refresh_token").
		Set("token_expires = EXCLUDED.token_expires").
		Set("updated_at = EXCLUDED.updated_at").
		Returning("*").
		Insert()
	return err
}
