package user

import (
	"context"
	"time"

	"github.com/fossadev/unbans/internal/db"
	"github.com/fossadev/unbans/internal/enums"
	"github.com/fossadev/unbans/internal/models"
	"github.com/fossadev/unbans/internal/twitchapi"
	"github.com/pkg/errors"
)

type UserFeature interface {
	RegisterTwitchUser(ctx context.Context, channelID int, user *twitchapi.User) (model *models.User, err error)
}

type userFeature struct {
	db *db.DB
}

func New(dbImpl *db.DB) UserFeature {
	return &userFeature{db: dbImpl}
}

func (f *userFeature) RegisterTwitchUser(ctx context.Context, channelID int, user *twitchapi.User) (*models.User, error) {
	model := &models.User{
		ChannelID:   channelID,
		Userlevel:   0,
		Login:       user.Login,
		DisplayName: user.DisplayName,
		Avatar:      user.ProfileImageURL,
		Provider:    enums.ProviderTwitch,
		ProviderID:  user.ID,
		Type:        user.Type,
		LastLogin:   time.Now().UTC(),
	}
	if err := f.db.Users.UpsertUser(ctx, model); err != nil {
		return nil, errors.Wrap(err, "failed to upsert user")
	}
	return model, nil
}
