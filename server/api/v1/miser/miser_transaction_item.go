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

type MiserTransactionItemApi struct{}

func (miserTransactionItemApi *MiserTransactionItemApi) CreateMiserTransactionItem(c *gin.Context) {
	var miserTransactionItem []miser.MiserTransactionItem
	err := c.ShouldBindJSON(&miserTransactionItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = miserTransactionItemService.CreateMiserTransactionItem(uid, miserTransactionItem)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (miserTransactionItemApi *MiserTransactionItemApi) DeleteMiserTransactionItem(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	err := miserTransactionItemService.DeleteMiserTransactionItem(uid, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (miserTransactionItemApi *MiserTransactionItemApi) DeleteMiserTransactionItemByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	uid := utils.GetUserID(c)
	err := miserTransactionItemService.DeleteMiserTransactionItemByIds(uid, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("批量删除成功", c)
}

func (miserTransactionItemApi *MiserTransactionItemApi) UpdateMiserTransactionItem(c *gin.Context) {
	var miserTransactionItem miser.MiserTransactionItem
	err := c.ShouldBindJSON(&miserTransactionItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = miserTransactionItemService.UpdateMiserTransactionItem(uid, miserTransactionItem)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

func (miserTransactionItemApi *MiserTransactionItemApi) FindMiserTransactionItem(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	miserTransactionItem, err := miserTransactionItemService.GetMiserTransactionItem(uid, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(miserTransactionItem, c)
}

func (miserTransactionItemApi *MiserTransactionItemApi) GetMiserTransactionItemList(c *gin.Context) {
	var pageInfo request.MiserTransactionItemSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	list, total, err := miserTransactionItemService.GetMiserTransactionItemInfoList(uid, pageInfo)
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

func (miserTransactionItemApi *MiserTransactionItemApi) ListItemDistinctNames(c *gin.Context) {
	uid := utils.GetUserID(c)
	list, err := miserTransactionItemService.ListItemDistinctNames(uid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(list, c)
}
