package router

import (
	"EduMall/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/edu/mall/login", handler.Login)

	authorized := r.Group("/")
	authorized.Use(handler.CheckLogin)
	eduRouter := authorized.Group("/edu/mall")
	{
		eduRouter.GET("/getUserInfo", handler.GetUserInfo)
	}
	return r
}
