package main

import "github.com/go-pg/migrations/v7"

func init() {
	migrations.MustRegister(func(db migrations.DB) error {
		_, err := db.Exec(`
			CREATE TYPE user_provider AS ENUM('twitch');

			CREATE TABLE users (
				id INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
				login VARCHAR(100) NOT NULL,
				display_name VARCHAR(120) NOT NULL,
				channel_id INTEGER NOT NULL,
				avatar VARCHAR NOT NULL,
				provider user_provider NOT NULL DEFAULT 'twitch',
				provider_id VARCHAR(64) NOT NULL,
				type VARCHAR(10) NOT NULL DEFAULT 'user',
				userlevel SMALLINT NOT NULL DEFAULT 0,
				last_login TIMESTAMPTZ NOT NULL,
				created_at TIMESTAMPTZ NOT NULL,
				updated_at TIMESTAMPTZ NOT NULL
			);

			CREATE UNIQUE INDEX users_provider_provider_id_idx ON users (provider, provider_id);
		`)
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(`
			DROP TABLE users;
			DROP TYPE user_provider;
		`)
		return err
	})
}