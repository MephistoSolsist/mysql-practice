package model

type Artist struct {
	ArtistId   int    `gorm:"PRIMARY_KEY" json:"artistId"`
	ArtistName string `json:"artistName"`
}

func (*Artist) TableName() string {
	return "artist"
}
