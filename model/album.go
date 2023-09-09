package model

type Album struct {
	AlbumId   int    `gorm:"PRIMARY_KEY" json:"artistId"`
	AlbumName string `json:"artistName"`
}

func (*Album) TableName() string {
	return "album"
}
