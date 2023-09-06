package initialize

import (
	"fmt"

	"github.com/MephistoSolsist/mysql-practice/config"
	"github.com/MephistoSolsist/mysql-practice/util"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func ReadConfig() config.ServerConfig {
	v := viper.New()

	v.SetConfigFile("config.yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("config read failed", err)
		util.Logger.Fatal(err)
	}
	serverConfig := config.ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		fmt.Println("config unmarshal failed", err)
		util.Logger.Fatal(err)
	}
	return serverConfig
}
