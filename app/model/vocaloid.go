package model

import (
	"github.com/naoina/genmai"
	"github.com/naoina/vocanew/db"
)

type Vocaloid struct {
	Id     int64  `db:"pk" json:"id"`
	Key    string `json:"key"`
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
	Order  int    `json:"order"`

	genmai.TimeStamp
}

func AllVocaloids() ([]*Vocaloid, error) {
	db := db.Get("default")
	var vocaloids []*Vocaloid
	return vocaloids, db.Select(&vocaloids, db.OrderBy("order", genmai.ASC))
}
