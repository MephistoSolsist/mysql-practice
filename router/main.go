package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `form:"username" json:"username" uri:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" binding:"required"`
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	UserGroup := r.Group("/user")
	{
		UserGroup.POST("login", UserRouterApp.Login)
		UserGroup.POST("register", UserRouterApp.Register)
	}

	return r
}

func LoginJSON(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if json.Username != "root" || json.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}
