package task

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/common/response"
	"github.com/springbear2020/self-hub/server/model/task"
	"github.com/springbear2020/self-hub/server/model/task/request"
	"github.com/springbear2020/self-hub/server/utils"
	"go.uber.org/zap"
)

type DailyTaskCompletionApi struct{}

func (taskCompletionApi *DailyTaskCompletionApi) CreateDailyTaskCompletion(c *gin.Context) {
	var taskCompletion task.DailyTaskCompletion
	err := c.ShouldBindJSON(&taskCompletion)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = dailyTaskCompletionService.CreateDailyTaskCompletion(uid, &taskCompletion)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (taskCompletionApi *DailyTaskCompletionApi) DeleteDailyTaskCompletion(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	err := dailyTaskCompletionService.DeleteDailyTaskCompletion(uid, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (taskCompletionApi *DailyTaskCompletionApi) DeleteDailyTaskCompletionByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	uid := utils.GetUserID(c)
	err := dailyTaskCompletionService.DeleteDailyTaskCompletionByIds(uid, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("批量删除成功", c)
}

func (taskCompletionApi *DailyTaskCompletionApi) UpdateDailyTaskCompletion(c *gin.Context) {
	var taskCompletion task.DailyTaskCompletion
	err := c.ShouldBindJSON(&taskCompletion)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = dailyTaskCompletionService.UpdateDailyTaskCompletion(uid, taskCompletion)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

func (taskCompletionApi *DailyTaskCompletionApi) FindDailyTaskCompletion(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	taskCompletion, err := dailyTaskCompletionService.GetDailyTaskCompletion(uid, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(taskCompletion, c)
}

func (taskCompletionApi *DailyTaskCompletionApi) GetDailyTaskCompletionList(c *gin.Context) {
	var pageInfo request.DailyTaskCompletionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	list, total, err := dailyTaskCompletionService.GetDailyTaskCompletionInfoList(uid, pageInfo)
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

func (taskCompletionApi *DailyTaskCompletionApi) GetDailyTaskStatList(c *gin.Context) {
	var pageInfo request.DailyTaskStat
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	list, total, err := dailyTaskCompletionService.GetDailyTaskStatInfoList(uid, pageInfo)
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
