package util_test

import (
	"fmt"
	"testing"

	"github.com/MephistoSolsist/mysql-practice/global"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Test_GetMusic(t *testing.T) {
	// var musicList []model.Music
	rows, err := global.DB.Table("music").Select("music.music_id,music.music_name,album.album_id,album.album_name,artist.artist_id,artist_name").Joins("left join artist on artist.artist_id = music.artist_id").Joins("left join album on album.album_id = music.album_id").Rows()
	fmt.Println(rows)
	fmt.Println(err)

	// global.DB.Joins("JOIN artist ON artist.artist_id = user.artist_id").Find(&musicList)
}
