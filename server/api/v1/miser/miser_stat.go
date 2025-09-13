package miser

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/common/response"
	"github.com/springbear2020/self-hub/server/model/miser/request"
	"github.com/springbear2020/self-hub/server/utils"
	"go.uber.org/zap"
)

type MiserStatApi struct {
}

func (a *MiserStatApi) GetCardStat(c *gin.Context) {
	var req request.MiserStat
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	result, err := miserStatService.GetCardStat(uid, &req)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

func (a *MiserStatApi) GetPieStat(c *gin.Context) {
	var req request.MiserStat
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	result, err := miserStatService.GetPieStat(uid, &req)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

func (a *MiserStatApi) GetLineStat(c *gin.Context) {
	var req request.MiserStat
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	result, err := miserStatService.GetLineStat(uid, &req)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

func (a *MiserStatApi) GetMonthStat(c *gin.Context) {
	var req request.MiserStat
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	result, err := miserStatService.GetMonthStat(uid, &req)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

func (a *MiserStatApi) GetYearStat(c *gin.Context) {
	var req request.MiserStat
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	result, err := miserStatService.GetYearStat(uid, &req)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

func (a *MiserStatApi) GetItemStat(c *gin.Context) {
	var req request.MiserStat
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	result, err := miserStatService.GetItemStat(uid, &req)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

func (a *MiserStatApi) GetCategoryStat(c *gin.Context) {
	var req request.MiserCategoryStat
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	result, err := miserStatService.GetCategoryStat(uid, &req)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

func (a *MiserStatApi) GetMonthTransactionStat(c *gin.Context) {
	var req request.MiserTransactionStat
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	result, err := miserStatService.GetMonthTransactionStat(uid, &req)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

func (a *MiserStatApi) GetCategoryItemStat(c *gin.Context) {
	var req request.MiserCategoryItemStat
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	result, err := miserStatService.GetCategoryItemStat(uid, &req)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}
