package model

type SongVocaloid struct {
	Id      int64  `db:"pk" column:"Id" json:"id"`
	VideoId string `column:"VideoId" json:"video_id"`
	Key     string `column:"Key" json:"key"`
}

func (sv *SongVocaloid) TableName() string {
	return "SongVocaloid"
}
