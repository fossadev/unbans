package twitch

import (
	"context"
	"net/http"
	"net/url"

	"github.com/fossadev/unbans/internal/encoder"
)

func (c *twitchController) redirect(ctx context.Context, w encoder.ResponseWriter, req *http.Request) {
	nextURL := ""
	if nextQuery := req.URL.Query().Get("next"); nextQuery != "" {
		if _, err := url.ParseRequestURI(nextQuery); err == nil {
			nextURL = nextQuery
		}
	}

	redirectURL, err := c.Features.TwitchAuth.GetAuthURL(nextURL)
	if err != nil {
		c.Log.WithRequest(req).Error("failed to generate auth url", err)
		w.InternalServerError("Failed to generate Twitch auth URL. Please try again later.")
		return
	}

	w.Redirect(redirectURL)
}
