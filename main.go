package main

import (
	"github.com/MephistoSolsist/mysql-practice/global"
	"github.com/MephistoSolsist/mysql-practice/initialize"
	"github.com/MephistoSolsist/mysql-practice/router"
	"github.com/MephistoSolsist/mysql-practice/util"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	global.CONFIG = initialize.ReadConfig()
	global.DB = initialize.InitDB()
	if global.DB != nil {
		defer global.DB.DB().Close()
	}
	r := router.SetupRouter()

	port := global.CONFIG.SystemConfig.Port
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)

	util.Logger.Info("Starting server at %v...", port)
}
