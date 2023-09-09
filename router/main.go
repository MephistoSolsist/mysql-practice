package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

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
		MusicGroup.DELETE("/:id", MusicRouterApp.Delete)
		MusicGroup.PUT("/", MusicRouterApp.Update)
		MusicGroup.GET("/:id", MusicRouterApp.GetMusicById)
	}

	return r
}
