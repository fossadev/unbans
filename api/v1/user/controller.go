package user

import "github.com/fossadev/unbans/internal/controller"

type userController struct {
	*controller.Controller
}

func New(parent *controller.Controller) {
	c := &userController{Controller: parent}

	c.Get("/user", c.getUser)
}
