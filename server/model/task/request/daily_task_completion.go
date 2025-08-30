package request

import (
	"github.com/springbear2020/self-hub/server/model/common/request"
	"time"
)

type DailyTaskCompletionSearch struct {
	TaskId     *int       `json:"taskId" form:"taskId"`
	FinishDate *time.Time `json:"finishDate" form:"finishDate"`
	request.PageInfo
}
