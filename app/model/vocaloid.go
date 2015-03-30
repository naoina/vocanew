package model

import (
	"github.com/naoina/genmai"
	"github.com/naoina/vocanew/db"
)

type Vocaloid struct {
	Id     int64  `db:"pk" column:"Id" json:"id"`
	Key    string `column:"Key" json:"key"`
	Name   string `column:"Name" json:"name"`
	NameEn string `column:"NameEn" json:"name_en"`
	Order  int    `column:"Order" json:"order"`

	genmai.TimeStamp
}

func AllVocaloids() ([]*Vocaloid, error) {
	db := db.Get("default")
	var vocaloids []*Vocaloid
	return vocaloids, db.Select(&vocaloids, db.OrderBy("Order", genmai.ASC))
}
