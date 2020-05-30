package twitch

import "github.com/fossadev/unbans/internal/controller"

type twitchController struct {
	*controller.Controller
}

func New(parent *controller.Controller) {
	c := &twitchController{Controller: parent}

	c.Get("/twitch/redirect", c.redirect)
	c.Get("/twitch/callback", c.callback)
}
