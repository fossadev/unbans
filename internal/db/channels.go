package db

import (
	"context"

	"github.com/fossadev/unbans/internal/models"
)

type Channels interface {
	UpsertChannel(ctx context.Context, model *models.Channel) (err error)
}
