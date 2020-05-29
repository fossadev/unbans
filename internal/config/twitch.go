package config

type TwitchConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

var TwitchScopes = []string{
	"moderation:read",
}
