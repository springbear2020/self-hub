package response

import "github.com/springbear2020/self-hub/server/model/task"

type DailyTaskStat struct {
	Task        task.DailyTask            `json:"task"`
	Completions []DailyTaskStatCompletion `json:"completions"`
}

type DailyTaskStatCompletion struct {
	FinishDate string `json:"finishDate"`
	CountValue int    `json:"countValue"`
}
