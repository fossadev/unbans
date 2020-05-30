package db

import (
	"context"

	"github.com/fossadev/unbans/internal/models"
)

type Users interface {
	UpsertUser(ctx context.Context, model *models.User) (err error)
}
