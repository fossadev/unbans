package postgres

import (
	"context"
	"time"

	"github.com/fossadev/unbans/internal/config"
	"github.com/fossadev/unbans/internal/db"
	"github.com/fossadev/unbans/internal/db/postgres/channels"
	"github.com/fossadev/unbans/internal/db/postgres/users"
	"github.com/go-pg/pg/v9"
)

type Postgres interface {
	Close(ctx context.Context) (err error)
	DB() *db.DB
}

type postgres struct {
	client *pg.DB
	db     *db.DB
}

func New(cfg *config.PostgresConfig) Postgres {
	conn := pg.Connect(&pg.Options{
		Addr:     cfg.Host,
		Network:  cfg.Network,
		User:     cfg.Username,
		Password: cfg.Password,
		Database: cfg.Database,

		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	dbImpl := &db.DB{
		Channels: channels.New(conn),
		Users:    users.New(conn),
	}

	return &postgres{client: conn, db: dbImpl}
}

func (p *postgres) DB() *db.DB {
	return p.db
}

func (p *postgres) Close(ctx context.Context) error {
	return p.client.WithContext(ctx).Close()
}
