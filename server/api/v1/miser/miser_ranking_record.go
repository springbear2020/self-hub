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

type MiserRankingRecordApi struct{}

func (miserRankingRecordApi *MiserRankingRecordApi) CreateMiserRankingRecord(c *gin.Context) {
	var miserRankingRecord miser.MiserRankingRecord
	err := c.ShouldBindJSON(&miserRankingRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = miserRankingRecordService.CreateMiserRankingRecord(uid, &miserRankingRecord)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (miserRankingRecordApi *MiserRankingRecordApi) DeleteMiserRankingRecord(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	err := miserRankingRecordService.DeleteMiserRankingRecord(uid, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (miserRankingRecordApi *MiserRankingRecordApi) DeleteMiserRankingRecordByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	uid := utils.GetUserID(c)
	err := miserRankingRecordService.DeleteMiserRankingRecordByIds(uid, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("批量删除成功", c)
}

func (miserRankingRecordApi *MiserRankingRecordApi) UpdateMiserRankingRecord(c *gin.Context) {
	var miserRankingRecord miser.MiserRankingRecord
	err := c.ShouldBindJSON(&miserRankingRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = miserRankingRecordService.UpdateMiserRankingRecord(uid, miserRankingRecord)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

func (miserRankingRecordApi *MiserRankingRecordApi) FindMiserRankingRecord(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	miserRankingRecord, err := miserRankingRecordService.GetMiserRankingRecord(uid, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(miserRankingRecord, c)
}

func (miserRankingRecordApi *MiserRankingRecordApi) GetMiserRankingRecordList(c *gin.Context) {
	var pageInfo request.MiserRankingRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	list, total, err := miserRankingRecordService.GetMiserRankingRecordInfoList(uid, pageInfo)
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
