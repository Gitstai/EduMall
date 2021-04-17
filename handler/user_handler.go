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
	user := GetUser(c)
	if user == nil || user.Id <= 0 {
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}
	//info, err := model.GetTUserInfo(&model.TUser{Id: user.Id})
	//if err != nil {
	//	logs.Logger.Errorf("func=%v, err=%v", "model.GetTUserInfo", err)
	//	ErrorHandler(c, config.ErrCodeErrBusinessException, "查询出错")
	//	return
	//}
	DataHandler(c, user)
	return
}

func Login(c *gin.Context) {
	//获取参数
	req := new(dto.LoginRequest)
	err := c.ShouldBindWith(req, binding.Form)
	if err != nil {
		logs.Logger.Errorf("req err:%v", err)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	//账号验证
	userInfo, err := model.GetTUserInfo(&model.TUser{Account: req.Account, Password: req.Password})
	if err != nil && err != gorm.ErrRecordNotFound {
		logs.Logger.Infof("系统错误, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrUserNotLogin, config.ErrMsgUserNotLogin)
		return
	} else if err == gorm.ErrRecordNotFound || userInfo == nil {
		logs.Logger.Infof("找不到该用户, account:%v, password:%v", req.Account, req.Password)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	// 账号验证通过，生成token
	tokenString, err := tools.GenerateToken(userInfo.Id)
	if err != nil {
		logs.Logger.Errorf("token生成错误, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgCodeErrBusinessException)
		return
	}

	logs.Logger.Infof("Login success! userInfo:%v", userInfo)
	DataHandler(c, tokenString)
}
