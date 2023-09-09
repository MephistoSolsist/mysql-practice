package router

import (
	"github.com/MephistoSolsist/mysql-practice/middleware"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `form:"username" json:"username" uri:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" binding:"required"`
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	UserGroup := r.Group("/user")
	{
		UserGroup.POST("login", UserRouterApp.Login)
		UserGroup.POST("register", UserRouterApp.Register)
	}

	MusicGroup := r.Group("/music")
	{
		MusicGroup.GET("/", MusicRouterApp.GetMusicList)
		MusicGroup.POST("/", MusicRouterApp.Upload)
		MusicGroup.DELETE("/", MusicRouterApp.Delete)
		MusicGroup.PUT("/", MusicRouterApp.Update)
		MusicGroup.GET("/:id",MusicRouterApp.GetMusicById)
	}

	return r
}
