package config

import "os"

type Config struct {
	Environment string
	WebHost     string

	Postgres *PostgresConfig
	Redis    *RedisConfig
	Twitch   *TwitchConfig
}

func New() *Config {
	return &Config{
		Environment: getEnvironment(env("UNBANS_ENVIRONMENT", "production")),
		WebHost:     env("UNBANS_WEB_HOST", ":3000"),

		Postgres: &PostgresConfig{},

		Redis: &RedisConfig{},

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
