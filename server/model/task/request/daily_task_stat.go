package request

import (
	"github.com/springbear2020/self-hub/server/model/common/request"
)

type DailyTaskStat struct {
	StartDate string `json:"startDate" form:"startDate"`
	EndDate   string `json:"endDate" form:"endDate"`
	request.PageInfo
}
