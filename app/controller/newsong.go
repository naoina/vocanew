package controller

import (
	"bytes"
	"html/template"
	"math"
	"strconv"

	"github.com/gorilla/feeds"
	"github.com/naoina/kocha"
	"github.com/naoina/vocanew/app/model"
)

var rssDescription = template.Must(template.New("description").Parse(`
<p><img src="{{ .ThumbnailUrl }}" alt="{{ .Title }}" width="94" height="70" /></p>
{{ .Description }}
<p><small><strong>{{ .FormatLength }}</strong> | <strong>{{ .FormatPostTime }}</strong> 投稿</small></p>
`))

type Newsong struct {
	*kocha.DefaultController
}

func (fe *Newsong) GET(c *kocha.Context) error {
	data := c.Data.(map[string]interface{})
	selectedVocaloids := data["SelectedVocaloids"].([]*model.Vocaloid)
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
	}
	latestUpdate, err := model.Songs.LatestUpdate()
	if err != nil {
		return c.RenderError(500, err, nil)
	}
	feed := &feeds.Feed{
		Title:       "VOCALOID新曲 - ぼかにゅー",
		Link:        &feeds.Link{Href: c.Request.URL.String(), Rel: "alternate"},
		Description: selectedVocaloids[0].Name + "の新曲",
		Author:      &feeds.Author{Name: "ぼかにゅー"},
		Updated:     latestUpdate,
	}
	var buf bytes.Buffer
	for _, v := range songs {
		buf.Reset()
		if err := rssDescription.Execute(&buf, v); err != nil {
			return c.RenderError(500, err, nil)
		}
		feed.Add(&feeds.Item{
			Title:       v.Title,
			Link:        &feeds.Link{Href: "http://www.nicovideo.jp/watch/" + v.VideoId, Rel: "alternate"},
			Description: buf.String(),
			Updated:     v.PostTime,
		})
	}
	var xml string
	if c.Params.Get("rss") == "2.0" {
		c.Response.ContentType = "application/rss+xml"
		xml, err = feed.ToRss()
	} else {
		c.Response.ContentType = "application/atom+xml"
		xml, err = feed.ToAtom()
	}
	if err != nil {
		return c.RenderError(500, err, nil)
	}
	return c.RenderText(xml)
}
