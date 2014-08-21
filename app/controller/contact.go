package controller

import (
	"github.com/naoina/kocha"
)

type Contact struct {
	*kocha.DefaultController
}

func (co *Contact) POST(c *kocha.Context) error {
	c.Response.ContentType = "application/json"
	if !c.Request.IsXHR() {
		return c.RenderError(404, nil, nil)
	}
	text := c.Params.Get("text")
	if text == "" {
		return c.RenderError(400, nil, map[string]interface{}{
			"success": 0,
			"errors":  []string{"内容が入力されていません。"},
		})
	}
	c.App.Event.Trigger("vocanew:send_contact_mail", text)
	return c.RenderJSON(map[string]interface{}{
		"success": 1,
	})
}
