package middleware

import (
	"github.com/naoina/kocha"
	"github.com/naoina/vocanew/app/model"
)

type Vocaloid struct {
}

func (m *Vocaloid) Process(app *kocha.Application, c *kocha.Context, next func() error) error {
	vocaloids, err := model.AllVocaloids()
	if err != nil {
		return err
	}
	selectedVocaloids := m.selectedVocaloids(c, vocaloids)
	keys := make([]string, len(selectedVocaloids))
	for i, v := range selectedVocaloids {
		keys[i] = v.Key
	}
	c.Data = map[string]interface{}{
		"Vocaloids":         vocaloids,
		"SelectedVocaloids": selectedVocaloids,
		"SelectedKeys":      keys,
	}
	return next()
}

func (m *Vocaloid) selectedVocaloids(c *kocha.Context, vocaloids []*model.Vocaloid) []*model.Vocaloid {
	var results []*model.Vocaloid
	for _, v := range vocaloids {
		if _, exists := c.Params.Values[v.Key]; exists {
			results = append(results, v)
		}
	}
	return results
}
