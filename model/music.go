package model

type Music struct {
	MusicId    int    `gorm:"PRIMARY_KEY" json:"musicId"`
	MusicName  string `json:"musicName"`
	ArtistId   int    `json:"artistId"`
	ArtistName string `json:"artistName"`
	AlbumId    int    `json:"albumId"`
	AlbumName  string `json:"albumName"`
}

type MusicDB struct {
	MusicId   int    `gorm:"PRIMARY_KEY" json:"musicId"`
	MusicName string `json:"musicName"`
	ArtistId  int    `json:"artistId"`
	AlbumId   int    `json:"albumId"`
}

func (*Music) TableName() string {
	return "music"
}

func (*MusicDB) TableName() string {
	return "music"
}
