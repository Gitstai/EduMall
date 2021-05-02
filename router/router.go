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

	r.POST("/edu/mall/login", handler.Login)       //done, test ok
	r.POST("/edu/mall/register", handler.Register) //done, test ok

	authorized := r.Group("/")
	authorized.Use(handler.CheckLogin)
	eduRouter := authorized.Group("/edu/mall")
	{
		eduRouter.GET("/logout", handler.Logout)                         //可不写 前端直接去除header
		eduRouter.GET("/getUserInfo", handler.GetUserInfo)               //done, test ok
		eduRouter.GET("/searchEduProducts", handler.SearchEduProducts)   //done, not test
		eduRouter.POST("/upsertEduProduct", handler.UpsertEduProduct)    //done, not test
		eduRouter.GET("/getProductDetail", handler.GetProductDetail)     //done, not test
		eduRouter.GET("/getProductEditInfo", handler.GetProductEditInfo) //done, not test
		eduRouter.GET("/getPurchaseRecords", handler.GetPurchaseRecords) //done, not test
		eduRouter.POST("/recharge", handler.Recharge)                    //done, not test
		eduRouter.GET("/purchase", handler.Purchase)                     //done, not test
		eduRouter.GET("/checkPurchased", handler.CheckPurchased)         //done, not test
	}
	return r
}
