package task

import "github.com/springbear2020/self-hub/server/service"

type ApiGroup struct {
	DailyTaskApi
	DailyTaskCompletionApi
}

var (
	dailyTaskService           = service.ServiceGroupApp.TaskServiceGroup.DailyTaskService
	dailyTaskCompletionService = service.ServiceGroupApp.TaskServiceGroup.DailyTaskCompletionService
)
