package handler

import (
	"EduMall/config"
	"EduMall/logs"
	"EduMall/model"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	id := int64(1)
	info, err := model.GetTUserInfo(&model.TUser{Id: id})
	if err != nil {
		logs.Error("func=%v, err=%v", "model.GetTUserInfo", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, "查询出错")
		return
	}
	logs.Info("success!!!!!")
	DataHandler(c, info)
	return
}
