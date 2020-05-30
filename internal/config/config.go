package config

import "os"

type Config struct {
	Environment string
	JWTSecret   string
	SiteURL     string
	WebHost     string

	Postgres *PostgresConfig
	Redis    *RedisConfig
	Twitch   *TwitchConfig
}

func New() *Config {
	return &Config{
		Environment: getEnvironment(env("UNBANS_ENVIRONMENT", "production")),
		JWTSecret:   env("UNBANS_JWT_SECRET", ""),
		SiteURL:     env("UNBANS_SITE_URL", ""),
		WebHost:     env("UNBANS_WEB_HOST", ":3000"),

		Postgres: &PostgresConfig{
			Host:     env("UNBANS_POSTGRES_HOST", "127.0.0.1:5432"),
			Network:  env("UNBANS_POSTGRES_NETWORK", "tcp"),
			Username: env("UNBANS_POSTGRES_USER", ""),
			Password: env("UNBANS_POSTGRES_PASSWORD", ""),
			Database: env("UNBANS_POSTGRES_DATABASE", "unbans"),
		},

		Redis: &RedisConfig{
			Network: env("UNBANS_REDIS_NETWORK", "tcp"),
			Host:    env("UNBANS_REDIS_HOST", "127.0.0.1:6379"),
		},

		Twitch: &TwitchConfig{
			ClientID:     env("UNBANS_TWITCH_CLIENT_ID", ""),
			ClientSecret: env("UNBANS_TWITCH_CLIENT_SECRET", ""),
			RedirectURL:  env("UNBANS_TWITCH_REDIRECT_URL", ""),
		},
	}
}

func env(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
