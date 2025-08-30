package task

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/middleware"
)

type DailyTaskRouter struct{}

func (s *DailyTaskRouter) InitDailyTaskRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	dailyTaskRouter := Router.Group("dailyTask").Use(middleware.OperationRecord())
	{
		dailyTaskRouter.POST("createDailyTask", dailyTaskApi.CreateDailyTask)
		dailyTaskRouter.DELETE("deleteDailyTask", dailyTaskApi.DeleteDailyTask)
		dailyTaskRouter.DELETE("deleteDailyTaskByIds", dailyTaskApi.DeleteDailyTaskByIds)
		dailyTaskRouter.PUT("updateDailyTask", dailyTaskApi.UpdateDailyTask)
	}

	dailyTaskRouterWithoutRecord := Router.Group("dailyTask")
	{
		dailyTaskRouterWithoutRecord.GET("findDailyTask", dailyTaskApi.FindDailyTask)
		dailyTaskRouterWithoutRecord.GET("getDailyTaskList", dailyTaskApi.GetDailyTaskList)
		dailyTaskRouterWithoutRecord.GET("listActiveTaskList", dailyTaskApi.ListActiveTaskList)
	}

	dailyTaskRouterWithoutAuth := PublicRouter.Group("dailyTask")
	{
		_ = dailyTaskRouterWithoutAuth
	}
}
