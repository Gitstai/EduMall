package handler

import (
	"EduMall/config"
	"EduMall/logs"
	"EduMall/model"
	"EduMall/tools"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CheckLogin(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		logs.Info("cookie中无token")
		ErrorHandler(c, config.ErrCodeErrUserNotLogin, config.ErrMsgUserNotLogin)
		c.Abort()
		return
	}
	id, isOk := tools.AuthCheck(token)
	if !isOk {
		logs.Info("token验证错误")
		ErrorHandler(c, config.ErrCodeErrUserNotLogin, config.ErrMsgUserNotLogin)
		c.Abort()
		return
	}
	userInfo, err := model.GetTUserInfo(&model.TUser{Id: id})
	if err != nil && err != gorm.ErrRecordNotFound {
		logs.Info("系统错误, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrUserNotLogin, config.ErrMsgUserNotLogin)
		c.Abort()
		return
	} else if err == gorm.ErrRecordNotFound {
		logs.Info("找不到该用户, id:%v", id)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		c.Abort()
		return
	}
	c.Set(config.UserInfo, userInfo)
	c.Next()
}