package service

import (
	"errors"

	"github.com/MephistoSolsist/mysql-practice/global"
	"github.com/MephistoSolsist/mysql-practice/model"
	"github.com/MephistoSolsist/mysql-practice/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserService struct{}

func (service *UserService) Register(user *model.User) error {
	var u model.User
	if !errors.Is(global.DB.Where("user_id = ?", user.UserId).First(&u).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已被注册")
	}
	user.Password = util.BcryptHash(user.Password)
	err := global.DB.Create(user).Error
	return err
}

func (service *UserService) Login(user *model.User) error {
	var u model.User
	err := global.DB.Where("user_id = ?", user.UserId).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("用户不存在")
	}
	if err == nil {
		if ok := util.BcryptCheck(user.Password, u.Password); !ok {
			return errors.New("密码错误")
		}
	}
	return err
}

var UserServiceApp = new(UserService)
