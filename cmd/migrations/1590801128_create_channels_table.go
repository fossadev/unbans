package main

import "github.com/go-pg/migrations/v7"

func init() {
	migrations.MustRegister(func(db migrations.DB) error {
		_, err := db.Exec(`
			CREATE TYPE channels_provider AS ENUM('twitch');

			CREATE TABLE channels (
				id INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
				login VARCHAR(100) NOT NULL,
				display_name VARCHAR(120) NOT NULL,
				avatar VARCHAR NOT NULL,
				provider channels_provider NOT NULL DEFAULT 'twitch',
				provider_id VARCHAR(64) NOT NULL,
				broadcaster_type VARCHAR(20) NOT NULL DEFAULT 'user',
				slug VARCHAR(104) NOT NULL,
				access_token VARCHAR(100) NOT NULL,
				refresh_token VARCHAR(100) NOT NULL,
				token_expires TIMESTAMPTZ NOT NULL,
				created_at TIMESTAMPTZ NOT NULL,
				updated_at TIMESTAMPTZ NOT NULL
			);

			CREATE UNIQUE INDEX channels_provider_provider_id_idx ON channels (provider, provider_id);
			CREATE UNIQUE INDEX channels_slug_lower_idx ON channels (lower(slug) varchar_pattern_ops);
		`)
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(`
			DROP TABLE channels;
			DROP TYPE channels_provider;
		`)
		return err
	})
}