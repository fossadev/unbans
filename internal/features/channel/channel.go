package channel

import (
	"context"
	"strconv"
	"strings"

	"github.com/fossadev/unbans/internal/db"
	"github.com/fossadev/unbans/internal/enums"
	"github.com/fossadev/unbans/internal/helpers"
	"github.com/fossadev/unbans/internal/logger"
	"github.com/fossadev/unbans/internal/models"
	"github.com/fossadev/unbans/internal/twitchapi"
	"github.com/go-pg/pg/v9"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

type ChannelFeature interface {
	RegisterTwitchChannel(ctx context.Context, token *oauth2.Token, user *twitchapi.User) (channel *models.Channel, err error)
}

type channelFeature struct {
	db  *db.DB
	log logger.Logger
}

func New(dbImpl *db.DB, log logger.Logger) ChannelFeature {
	return &channelFeature{
		db:  dbImpl,
		log: log,
	}
}

func (f *channelFeature) RegisterTwitchChannel(ctx context.Context, token *oauth2.Token, user *twitchapi.User) (*models.Channel, error) {
	for i := 0; i < 10; i++ {
		channel := &models.Channel{
			Login:           user.Login,
			DisplayName:     user.DisplayName,
			Avatar:          user.ProfileImageURL,
			Provider:        enums.ProviderTwitch,
			ProviderID:      user.ID,
			BroadcasterType: user.BroadcasterType,
			Slug:            strings.ToLower(user.Login),
			AccessToken:     token.AccessToken,
			RefreshToken:    token.RefreshToken,
			TokenExpires:    token.Expiry,
		}

		if i != 0 {
			channel.Slug += "-" + strconv.Itoa(helpers.MathRandInt(1000, 9999))
		}

		if err := f.db.Channels.UpsertChannel(ctx, channel); err != nil {
			if pgErr, ok := err.(pg.Error); ok && pgErr.IntegrityViolation() {
				continue
			}
			return nil, errors.Wrap(err, "failed to upsert channel")
		}

		return channel, nil
	}

	return nil, errors.New("failed to generate unique slug for channel")
}
