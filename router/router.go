package router

import (
	"EduMall/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(handler.Cors())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/edu/mall/login", handler.Login)
	r.POST("/edu/mall/register", handler.Register)

	authorized := r.Group("/")
	authorized.Use(handler.CheckLogin)
	eduRouter := authorized.Group("/edu/mall")
	{
		eduRouter.GET("/logout", handler.Logout) //可不写 前端直接去除header
		eduRouter.GET("/getUserInfo", handler.GetUserInfo)
		eduRouter.GET("/searchEduProducts", handler.SearchEduProducts)
		eduRouter.POST("/upsertEduProduct", handler.UpsertEduProduct)
		eduRouter.GET("/getProductDetail", handler.GetProductDetail)
		eduRouter.GET("/getProductEditInfo", handler.GetProductEditInfo)
		eduRouter.GET("/getPurchaseRecords", handler.GetPurchaseRecords)
		eduRouter.GET("/recharge", handler.Recharge)
		eduRouter.GET("/purchase", handler.Purchase)
		eduRouter.GET("/checkPurchased", handler.CheckPurchased)
	}
	return r
}
