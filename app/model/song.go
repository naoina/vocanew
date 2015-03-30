package model

import (
	"fmt"
	"time"

	"github.com/naoina/genmai"
	"github.com/naoina/vocanew/db"
)

type Song struct {
	Id           int64     `db:"pk" column:"Id" json:"id"`
	VideoId      string    `column:"VideoId" json:"video_id"`
	Title        string    `column:"Title" json:"title"`
	Description  string    `column:"Description" json:"description"`
	Length       int64     `column:"Length" json:"length"`
	ThumbnailUrl string    `column:"ThumbnailUrl" json:"thumbnail_url"`
	PostTime     time.Time `column:"PostTime" json:"post_time"`
	Tags         string    `column:"Tags" json:"tags"`
	AddTime      time.Time `column:"AddTime" json:"add_time"`

	genmai.TimeStamp
}

var Songs = &Song{}

func (m *Song) Count(vocaloids ...*Vocaloid) (int64, error) {
	db := db.Get("default")
	var count int64
	if len(vocaloids) < 1 {
		return count, db.Select(&count, db.Count("Id"), db.From(&Song{}))
	}
	where := db.Where("Key", "=", vocaloids[0].Key)
	for _, v := range vocaloids[1:] {
		where = where.Or(db.Where("Key", "=", v.Key))
	}
	return count, db.Select(&count, db.Count("Id"), db.From(&SongVocaloid{}), where)
}

func (m *Song) LatestUpdate() (time.Time, error) {
	db := db.Get("default")
	var songs []Song
	if err := db.Select(&songs, db.OrderBy("Id", genmai.DESC).Limit(1)); err != nil {
		return time.Time{}, err
	}
	return songs[0].AddTime, nil
}

func (m *Song) FindByVocaloid(vocaloids []*Vocaloid, offset, limit int) ([]*Song, error) {
	var songs []*Song
	db := db.Get("default")
	if len(vocaloids) == 0 {
		return songs, db.Select(&songs, db.Distinct("*"), db.OrderBy("Id", genmai.DESC).Offset(offset).Limit(limit))
	}
	keys := make([]interface{}, len(vocaloids))
	for i, v := range vocaloids {
		keys[i] = v.Key
	}
	var songVocaloids []SongVocaloid
	if err := db.Select(&songVocaloids, db.Distinct("VideoId"), db.Where("Key").In(keys...), db.OrderBy("Id", genmai.DESC).Offset(offset).Limit(limit)); err != nil {
		return nil, err
	}
	keys = keys[:0]
	for _, sv := range songVocaloids {
		keys = append(keys, sv.VideoId)
	}
	return songs, db.Select(&songs, db.Distinct("*"), db.Where("VideoId").In(keys...), db.OrderBy("Id", genmai.DESC))
}

func (m *Song) FormatLength() string {
	return fmt.Sprintf("%d:%02d", m.Length/60, m.Length%60)
}

func (m *Song) FormatPostTime() string {
	return m.PostTime.Format("2006年01月02日 15:04")
}

func (m *Song) Vocaloids() ([]Vocaloid, error) {
	db := db.Get("default")
	var vocaloids []Vocaloid
	if err := db.Select(&vocaloids, db.Join(&SongVocaloid{}).On("Key").Where(&SongVocaloid{}, "VideoId", "=", m.VideoId)); err != nil {
		return nil, err
	}
	return vocaloids, nil
}
