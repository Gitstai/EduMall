package handler

import (
	"EduMall/config"
	"EduMall/dto"
	"EduMall/logs"
	"EduMall/model"
	"EduMall/tools"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
)

func GetUserInfo(c *gin.Context) {
	id := int64(1)
	info, err := model.GetTUserInfo(&model.TUser{Id: id})
	if err != nil {
		logs.Error("func=%v, err=%v", "model.GetTUserInfo", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, "查询出错")
		return
	}
	DataHandler(c, info)
	return
}

func Login(c *gin.Context) {
	//获取参数
	req := new(dto.LoginRequest)
	err := c.ShouldBindWith(req, binding.Form)
	if err != nil {
		logs.Error("req err:%v", err)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	//账号验证
	userInfo, err := model.GetTUserInfo(&model.TUser{Account: req.Account, Password: req.Password})
	if err != nil && err != gorm.ErrRecordNotFound {
		logs.Info("系统错误, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrUserNotLogin, config.ErrMsgUserNotLogin)
		return
	} else if err == gorm.ErrRecordNotFound {
		logs.Info("找不到该用户, account:%v, password:%v", req.Account, req.Password)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	// 账号验证通过，生成token
	tokenString, err := tools.GenerateToken(1)
	if err != nil {
		logs.Error("token生成错误, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgCodeErrBusinessException)
		return
	}

	logs.Info("Login success! userInfo:%v", userInfo)
	DataHandler(c, tokenString)
}
