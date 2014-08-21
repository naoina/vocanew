package model

type SongVocaloid struct {
	Id      int64  `db:"pk" json:"id"`
	VideoId string `json:"video_id"`
	Key     string `json:"key"`
}
