package model

type Music struct {
	MusicId    int    `json:"musicId"`
	MusicName  string `json:"musicName"`
	ArtistId   int    `json:"artistId"`
	ArtistName string `json:"artistName"`
	AlbumId    int    `json:"albumId"`
	AlbumName  string `json:"albumName"`
}

type MusicDB struct {
	MusicId   int `gorm:"PRIMARY_KEY"`
	MusicName string
	ArtistId  int
	AlbumId   int
}

func (*Music) TableName() string {
	return "music"
}

func (*MusicDB) TableName() string {
	return "music"
}
