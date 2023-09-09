package req

type UploadMusicReq struct {
	MusicName string `json:"musicName"`
	AlbumId   int    `json:"albumId"`
	ArtistId  int    `json:"artistId"`
}

type DeleteMusicReq struct {
	MusicId int `json:"musicId"`
}

func (*UploadMusicReq) TableName() string {
	return "music"
}
