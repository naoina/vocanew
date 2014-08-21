package controller

import (
	"github.com/naoina/kocha"
)

type About struct {
	*kocha.DefaultController
}

func (ab *About) GET(c *kocha.Context) error {
	return c.Render(nil)
}
