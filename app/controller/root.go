package controller

import (
	"math"
	"strconv"

	"github.com/naoina/kocha"
	"github.com/naoina/vocanew/app/lib/pagination"
	"github.com/naoina/vocanew/app/model"
)

const (
	itemsPerPage        = 10
	maxPaginationLength = 10
	paginationThreshold = 5
)

type Root struct {
	*kocha.DefaultController
}

func (r *Root) GET(c *kocha.Context) error {
	data := c.Data.(map[string]interface{})
	selectedVocaloids := data["SelectedVocaloids"].([]*model.Vocaloid)
	count, err := model.Songs.Count(selectedVocaloids...)
	if err != nil {
		return c.RenderError(500, err, nil)
	}
	p, err := strconv.Atoi(c.Params.Get("p"))
	if err != nil {
		p = 1
	}
	offset := int(math.Max(itemsPerPage, 0) * math.Max(float64(p-1), 0))
	songs, err := model.Songs.FindByVocaloid(selectedVocaloids, offset, itemsPerPage)
	if err != nil {
		return c.RenderError(500, err, nil)
	}
	if len(selectedVocaloids) == 0 {
		selectedVocaloids = append(selectedVocaloids, &model.Vocaloid{Key: "vocaloid", Name: "VOCALOID"})
		data["SelectedVocaloids"] = selectedVocaloids
	}
	paginationLength := int(count / itemsPerPage)
	if count%itemsPerPage > 0 {
		paginationLength++
	}
	plength := paginationLength
	if plength < 1 {
		plength = 1
	}
	if plength > maxPaginationLength {
		plength = maxPaginationLength
	}
	paginator := pagination.NewPaginator(count, itemsPerPage, plength, paginationThreshold)
	page := paginator.Page(p)
	pages := paginator.Pages(p)
	return c.Render(map[string]interface{}{
		"Songs":            songs,
		"SelectedVocaloid": selectedVocaloids[0],
		"Pages":            pages,
		"Page":             page,
	})
}
