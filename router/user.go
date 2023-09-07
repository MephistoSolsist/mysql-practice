package router

import (
	"github.com/MephistoSolsist/mysql-practice/model"
	"github.com/MephistoSolsist/mysql-practice/service"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (UserRouter) Register(c *gin.Context) {
	var u model.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		return
	}
	err = service.UserServiceApp.Register(&u)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, u)

}

func (UserRouter) Login(c *gin.Context) {
	var u model.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		return
	}
	err = service.UserServiceApp.Login(&u)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, u)

}

var UserRouterApp = new(UserRouter)
