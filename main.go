package main

import (
	"fmt"

	"github.com/MephistoSolsist/mysql-practice/global"
	"github.com/MephistoSolsist/mysql-practice/initialize"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Student struct {
	SId      string `gorm:"column:SId"`
	GId      string `gorm:"column:GId"`
	Name     string `gorm:"column:SName"`
	Sex      string `gorm:"column:SSexy"`
	Birthday string `gorm:"column:SBdate"`
	Phone    string `gorm:"column:STele"`
}

func (s *Student) TableName() string {
	return "student"
}
func main() {
	global.CONFIG = initialize.ReadConfig()
	global.DB = initialize.InitDB()
	if global.DB != nil {
		defer global.DB.DB().Close()
	}
	var s = new(Student)
	global.DB.First(s)
	fmt.Println(s.Name)
}
