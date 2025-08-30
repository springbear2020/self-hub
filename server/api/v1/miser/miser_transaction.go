package miser

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/common/response"
	"github.com/springbear2020/self-hub/server/model/miser"
	"github.com/springbear2020/self-hub/server/model/miser/request"
	"github.com/springbear2020/self-hub/server/utils"
	"go.uber.org/zap"
)

type MiserTransactionApi struct{}

func (miserTransactionApi *MiserTransactionApi) CreateMiserTransaction(c *gin.Context) {
	var miserTransaction []miser.MiserTransaction
	err := c.ShouldBindJSON(&miserTransaction)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = miserTransactionService.CreateMiserTransaction(uid, miserTransaction)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (miserTransactionApi *MiserTransactionApi) DeleteMiserTransaction(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	err := miserTransactionService.DeleteMiserTransaction(uid, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (miserTransactionApi *MiserTransactionApi) DeleteMiserTransactionByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	uid := utils.GetUserID(c)
	err := miserTransactionService.DeleteMiserTransactionByIds(uid, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("批量删除成功", c)
}

func (miserTransactionApi *MiserTransactionApi) UpdateMiserTransaction(c *gin.Context) {
	var miserTransaction miser.MiserTransaction
	err := c.ShouldBindJSON(&miserTransaction)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = miserTransactionService.UpdateMiserTransaction(uid, miserTransaction)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

func (miserTransactionApi *MiserTransactionApi) FindMiserTransaction(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	miserTransaction, err := miserTransactionService.GetMiserTransaction(uid, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(miserTransaction, c)
}

func (miserTransactionApi *MiserTransactionApi) GetMiserTransactionList(c *gin.Context) {
	var pageInfo request.MiserTransactionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	list, total, err := miserTransactionService.GetMiserTransactionInfoList(uid, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
