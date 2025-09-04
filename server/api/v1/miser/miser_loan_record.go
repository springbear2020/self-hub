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

type MiserLoanRecordApi struct{}

func (miserLoanRecordApi *MiserLoanRecordApi) CreateMiserLoanRecord(c *gin.Context) {
	var miserLoanRecord miser.MiserLoanRecord
	err := c.ShouldBindJSON(&miserLoanRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = miserLoanRecordService.CreateMiserLoanRecord(uid, &miserLoanRecord)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (miserLoanRecordApi *MiserLoanRecordApi) DeleteMiserLoanRecord(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	err := miserLoanRecordService.DeleteMiserLoanRecord(uid, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (miserLoanRecordApi *MiserLoanRecordApi) DeleteMiserLoanRecordByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	uid := utils.GetUserID(c)
	err := miserLoanRecordService.DeleteMiserLoanRecordByIds(uid, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("批量删除成功", c)
}

func (miserLoanRecordApi *MiserLoanRecordApi) UpdateMiserLoanRecord(c *gin.Context) {
	var miserLoanRecord miser.MiserLoanRecord
	err := c.ShouldBindJSON(&miserLoanRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = miserLoanRecordService.UpdateMiserLoanRecord(uid, miserLoanRecord)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

func (miserLoanRecordApi *MiserLoanRecordApi) FindMiserLoanRecord(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	miserLoanRecord, err := miserLoanRecordService.GetMiserLoanRecord(uid, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(miserLoanRecord, c)
}

func (miserLoanRecordApi *MiserLoanRecordApi) GetMiserLoanRecordList(c *gin.Context) {
	var pageInfo request.MiserLoanRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	list, total, err := miserLoanRecordService.GetMiserLoanRecordInfoList(uid, pageInfo)
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

func (miserLoanRecordApi *MiserLoanRecordApi) ListMiserLoanNameList(c *gin.Context) {
	uid := utils.GetUserID(c)
	list, err := miserLoanRecordService.ListMiserLoanNameList(uid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(list, c)
}

func (miserLoanRecordApi *MiserLoanRecordApi) GetMiserLoanStatData(c *gin.Context) {
	uid := utils.GetUserID(c)
	data, err := miserLoanRecordService.GetMiserLoanStatData(uid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}
