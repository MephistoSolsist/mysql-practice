package service

import (
	"github.com/MephistoSolsist/mysql-practice/global"
	"github.com/MephistoSolsist/mysql-practice/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MusicService struct{}

func (*MusicService) GetMusicList(list *[]model.Music) error {
	err := global.DB.Raw("select music_name,music.music_id,album_name,album.album_id,artist.artist_id,artist_name from music,album,artist where ifnull(music.artist_id,0) = artist.artist_id and IFNULL(music.album_id,0)=album.album_id").Scan(&list)
	if err != nil {
		return err.Error
	}
	return nil
}

func (*MusicService) GetMusicById(id int) (model.Music, error) {
	var m model.Music
	err := global.DB.Raw("select music_name,music.music_id,album_name,album.album_id,artist.artist_id,artist_name from music,album,artist where ifnull(music.artist_id,0) = artist.artist_id and IFNULL(music.album_id,0)=album.album_id and music_id = ?", id).Scan(&m)
	if err != nil {
		return m, err.Error
	}
	return m, nil
}

func (*MusicService) DeleteMusic(m *model.MusicDB) error {
	err := global.DB.Delete(&m)
	if err != nil {
		return err.Error
	}
	return nil
}

func (*MusicService) UploadMusic(m *model.MusicDB) error {
	err := global.DB.Create(&m)
	if err != nil {
		return err.Error
	}
	return nil
}

func (*MusicService) UpdateMusic(m *model.MusicDB) error {
	global.DB.Save(m)
	return nil
}

var MusicServiceApp = new(MusicService)
