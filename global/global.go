package global

import (
	"github.com/jinzhu/gorm"
	"github.com/MephistoSolsist/mysql-practice/config"
)

var (
	DB     *gorm.DB
	CONFIG config.ServerConfig
)
