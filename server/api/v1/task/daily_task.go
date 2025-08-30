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

type DailyTaskApi struct{}

func (dailyTaskApi *DailyTaskApi) CreateDailyTask(c *gin.Context) {
	var dailyTask task.DailyTask
	err := c.ShouldBindJSON(&dailyTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = dailyTaskService.CreateDailyTask(uid, &dailyTask)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (dailyTaskApi *DailyTaskApi) DeleteDailyTask(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	err := dailyTaskService.DeleteDailyTask(uid, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (dailyTaskApi *DailyTaskApi) DeleteDailyTaskByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	uid := utils.GetUserID(c)
	err := dailyTaskService.DeleteDailyTaskByIds(uid, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("批量删除成功", c)
}

func (dailyTaskApi *DailyTaskApi) UpdateDailyTask(c *gin.Context) {
	var dailyTask task.DailyTask
	err := c.ShouldBindJSON(&dailyTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = dailyTaskService.UpdateDailyTask(uid, dailyTask)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

func (dailyTaskApi *DailyTaskApi) FindDailyTask(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	dailyTask, err := dailyTaskService.GetDailyTask(uid, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(dailyTask, c)
}

func (dailyTaskApi *DailyTaskApi) GetDailyTaskList(c *gin.Context) {
	var pageInfo request.DailyTaskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	list, total, err := dailyTaskService.GetDailyTaskInfoList(uid, pageInfo)
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

func (dailyTaskApi *DailyTaskApi) ListActiveTaskList(c *gin.Context) {
	uid := utils.GetUserID(c)
	dailyTasks, err := dailyTaskService.ListActiveTaskList(uid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(dailyTasks, c)
}
