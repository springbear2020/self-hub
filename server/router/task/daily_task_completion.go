package task

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/middleware"
)

type DailyTaskCompletionRouter struct{}

func (s *DailyTaskCompletionRouter) InitDailyTaskCompletionRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	taskCompletionRouter := Router.Group("dailyTaskCompletion").Use(middleware.OperationRecord())
	{
		taskCompletionRouter.POST("createDailyTaskCompletion", dailyTaskCompletionApi.CreateDailyTaskCompletion)
		taskCompletionRouter.DELETE("deleteDailyTaskCompletion", dailyTaskCompletionApi.DeleteDailyTaskCompletion)
		taskCompletionRouter.DELETE("deleteDailyTaskCompletionByIds", dailyTaskCompletionApi.DeleteDailyTaskCompletionByIds)
		taskCompletionRouter.PUT("updateDailyTaskCompletion", dailyTaskCompletionApi.UpdateDailyTaskCompletion)
	}

	taskCompletionRouterWithoutRecord := Router.Group("dailyTaskCompletion")
	{
		taskCompletionRouterWithoutRecord.GET("findDailyTaskCompletion", dailyTaskCompletionApi.FindDailyTaskCompletion)
		taskCompletionRouterWithoutRecord.GET("getDailyTaskCompletionList", dailyTaskCompletionApi.GetDailyTaskCompletionList)
		taskCompletionRouterWithoutRecord.GET("getDailyTaskStatList", dailyTaskCompletionApi.GetDailyTaskStatList)
	}

	taskCompletionRouterWithoutAuth := PublicRouter.Group("dailyTaskCompletion")
	{
		_ = taskCompletionRouterWithoutAuth
	}
}
