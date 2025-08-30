package task

import api "github.com/springbear2020/self-hub/server/api/v1"

type RouterGroup struct {
	DailyTaskRouter
	DailyTaskCompletionRouter
}

var (
	dailyTaskApi           = api.ApiGroupApp.TaskApiGroup.DailyTaskApi
	dailyTaskCompletionApi = api.ApiGroupApp.TaskApiGroup.DailyTaskCompletionApi
)
