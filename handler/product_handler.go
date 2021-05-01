package handler

import (
	"EduMall/config"
	"EduMall/dto"
	"EduMall/logs"
	"EduMall/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strings"
)

func SearchEduProducts(c *gin.Context) {
	user := GetUser(c)
	if user == nil || user.Id <= 0 {
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	req := new(dto.SearchEduProductsReq)
	err := c.ShouldBindQuery(req)
	if err != nil {
		logs.Logger.Infof("req err:%v", req)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	prods, total, err := model.QueryTProduct(&model.TProduct{ProviderName: req.Provider, ProductType: int8(req.ProductType), Keywords: req.Keywords}, req.PageNum, req.PageSize)
	if err != nil {
		logs.Logger.Errorf("func:%v, err:%v", "model.QueryTProduct", err)
		ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgCodeErrBusinessException)
		return
	}

	res := make([]*dto.EduProduct, 0, len(prods))
	for _, p := range prods {
		tmp := new(dto.EduProduct)
		tmp.ProductId = p.ProviderId
		tmp.Provider = p.ProviderName
		tmp.ProductType = int32(p.ProductType)
		tmp.ProductDesc = p.DescText
		tmp.ProductName = p.Name
		tmp.ProductImg = p.DescImg
		tmp.CreatedTime = p.CreateTime.Unix()
		tmp.SaleVolume = p.SaleVolume
		tmp.Inventory = p.Inventory

		res = append(res, tmp)
	}

	DataHandlerWithTotal(c, res, total)
}

func UpsertEduProduct(c *gin.Context) {
	user := GetUser(c)
	if user == nil || user.Id <= 0 {
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	req := new(dto.UpsertEduProductReq)
	err := c.ShouldBindWith(req, binding.Form)
	if err != nil {
		logs.Logger.Infof("req err:%v", req)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	//验证传过来的userid是否和登录的userid一致
	if req.UserId == 0 || req.UserId != user.Id {
		logs.Logger.Infof("userId is not equal, req.UserId:%v, userId:%v", req.UserId, user.Id)
		ErrorHandler(c, config.ErrCodeErrREQParamInvalid, config.ErrMsgREQParamInvalid)
		return
	}

	if req.Id != 0 { //编辑产品
		//先查询是否有该条记录，并进行providerId和userId的验证
		//TODO 暂时不查

		//更新产品信息
		if err := model.UpdateTProduct(&model.TProduct{
			Id:            req.Id,
			Name:          req.ProductName,
			ProductType:   int8(req.ProductType),
			Status:        int8(req.Status),
			Price:         req.Price,
			AfterSaleText: req.AfterSaleText,
			DescText:      req.ProductDesc,
			DescImg:       strings.Join(req.BannerImgs, ";"),
		}); err != nil {
			logs.Logger.Errorf("func:%v, err:%v", "model.UpdateTProduct", err)
			ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgCodeErrBusinessException)
			return
		}
	} else { //创建产品
		p, err := model.InsertTProduct(&model.TProduct{
			ProviderId:    user.Id,
			ProviderName:  user.Nickname,
			Name:          req.ProductName,
			ProductType:   int8(req.ProductType),
			Status:        int8(req.Status),
			Price:         req.Price,
			AfterSaleText: req.AfterSaleText,
			DescText:      req.ProductDesc,
			DescImg:       strings.Join(req.BannerImgs, ";"),
		})
		if err != nil {
			logs.Logger.Errorf("func:%v, err:%v", "model.InsertTProduct", err)
			ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgCodeErrBusinessException)
			return
		}
		req.Id = p.Id
	}

	if req.FileId != 0 { //编辑产品对应的文件
		//先查询是否有该条记录
		//TODO 暂时不查

		//更新文件信息
		if err := model.UpdateTProductFile(&model.TProductFile{
			Id:       req.FileId,
			FileType: int8(req.FileType),
			FileName: req.FileName,
			Url:      req.FileUrl,
		}); err != nil {
			logs.Logger.Errorf("func:%v, err:%v", "model.UpdateTProductFile", err)
			ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgCodeErrBusinessException)
			return
		}
	} else { //创建产品对应的文件
		if _, err := model.InsertTProductFile(&model.TProductFile{
			FileType:  int8(req.FileType),
			FileName:  req.FileName,
			Url:       req.FileUrl,
			ProductID: req.Id,
		}); err != nil {
			logs.Logger.Errorf("func:%v, err:%v", "model.InsertTProductFile", err)
			ErrorHandler(c, config.ErrCodeErrBusinessException, config.ErrMsgCodeErrBusinessException)
			return
		}
	}

	DataHandler(c, nil)
}

func GetProductDetail(c *gin.Context) {

}

func GetProductEditInfo(c *gin.Context) {

}
