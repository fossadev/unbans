package twitch

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/fossadev/unbans/internal/enums"
	"github.com/fossadev/unbans/internal/twitchapi"

	"github.com/fossadev/unbans/internal/encoder"
)

func (c *twitchController) callback(ctx context.Context, w encoder.ResponseWriter, req *http.Request) {
	code := strings.TrimSpace(req.URL.Query().Get("code"))
	state := strings.TrimSpace(req.URL.Query().Get("state"))
	if code == "" || state == "" {
		w.BadRequest("code or state parameter missing")
		return
	}

	token, nextURL, err := c.Features.TwitchAuth.ExchangeCode(ctx, code, state)
	if err != nil {
		c.Log.WithRequest(req).Error("failed to verify exchange code", err)
		w.BadRequest("failed to validate oauth code")
		return
	}

	users, err := c.TwitchAPI.GetUsers(&twitchapi.GetUsersRequest{Token: token.AccessToken})
	if err != nil || len(users) != 1 {
		if err != nil {
			c.Log.WithRequest(req).Error("failed to get user from helix", err)
		}
		w.BadRequest("Twitch API failed to find your user. Try authing again, Twitch messed up.")
		return
	}

	twitchUser := users[0]

	channel, err := c.Features.Channel.RegisterTwitchChannel(ctx, token, twitchUser)
	if err != nil {
		c.Log.WithRequest(req).Error("failed to register channel", err)
		w.InternalServerError("Failed to register channel, please try again later.")
		return
	}

	user, err := c.Features.User.RegisterTwitchUser(ctx, channel.ID, twitchUser)
	if err != nil {
		c.Log.WithRequest(req).Error("failed to register user", err)
		w.InternalServerError("Failed to register user, please try again later.")
		return
	}

	jwtToken, err := c.Features.Token.CreateJWT(user.ID, channel.ID, enums.ProviderTwitch, twitchUser.ID)
	if err != nil {
		c.Log.WithRequest(req).Error("failed to generate jwt", err)
		w.InternalServerError("Failed to generate user authorization token, please try again later.")
		return
	}

	w.SetCookie(&http.Cookie{
		Name:   "unbanssess",
		Value:  jwtToken,
		Path:   "/",
		Domain: ".",
		Secure: false,
	})

	if nextURL == "" {
		w.Redirect(fmt.Sprintf("%s/%s/dashboard", c.Config.SiteURL, url.PathEscape(channel.Slug)))
		return
	}

	w.Success(nextURL)
}
