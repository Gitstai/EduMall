package handler

import (
	"EduMall/config"
	"EduMall/dto"
	"EduMall/logs"
	"EduMall/model"
	"EduMall/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"math/rand"
	"time"
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
	if err != nil {
		logs.Logger.Infof("系统错误, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrUserNotLogin, config.ErrMsgUserNotLogin)
		return
	} else if len(userInfo) != 1 {
		logs.Logger.Infof("找不到该用户, account:%v, password:%v, len(userInfo)=%v", req.Account, req.Password, len(userInfo))
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgUserWrongUser)
		return
	}

	// 账号验证通过，生成token
	tokenString, err := tools.GenerateToken(userInfo[0].Id)
	if err != nil {
		logs.Logger.Errorf("token生成错误, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgCodeErrBusinessException)
		return
	}

	logs.Logger.Infof("Login success! userInfo:%v", tools.ToJson(userInfo))
	c.SetCookie("token", tokenString, int(time.Hour), "/", "", false, false)
	DataHandler(c, dto.LoginResponse{UserId: userInfo[0].Id})
}

func Register(c *gin.Context) {
	//获取参数
	req := new(dto.RegisterRequest)
	err := c.ShouldBindWith(req, binding.Form)
	if err != nil {
		logs.Logger.Errorf("req err:%v", err)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	if req.Password == "" || req.Account == "" {
		logs.Logger.Errorf("req err:%v", err)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	//账号是否重复
	userInfo, err := model.GetTUserInfo(&model.TUser{Account: req.Account})
	if err != nil {
		logs.Logger.Infof("系统错误, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrUserNotLogin, config.ErrMsgUserNotLogin)
		return
	} else if len(userInfo) != 0 {
		logs.Logger.Infof("账号重复, account:%v, password:%v", req.Account, req.Password)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	//建立账号
	if req.Nickname == nil || *req.Nickname == "" {
		str := fmt.Sprintf("昵称%d", rand.Uint64())
		req.Nickname = &str
	}
	user, err := model.InsertTUser(&model.TUser{Account: req.Account, Password: req.Password, Nickname: *req.Nickname})
	if err != nil {
		logs.Logger.Infof("系统错误, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, "注册失败，请稍后重试")
		return
	}

	// 账号注册成功，生成token
	tokenString, err := tools.GenerateToken(user.Id)
	if err != nil {
		logs.Logger.Errorf("token生成错误, err:%v", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgCodeErrBusinessException)
		return
	}

	logs.Logger.Infof("Register success! userInfo:%v", tools.ToJson(user))
	c.SetCookie("token", tokenString, int(time.Hour), "/", "", false, false)
	DataHandler(c, dto.RegisterResponse{UserId: user.Id})
}
