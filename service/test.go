package service

import (
	"fmt"

	"github.com/MephistoSolsist/mysql-practice/global"
	"github.com/MephistoSolsist/mysql-practice/model"
)

func Test() {
	var user = new(model.User)
	global.DB.First(&user)
	fmt.Printf("%#v", user)
}
