package request

import (
	"github.com/springbear2020/self-hub/server/model/common/request"
)

type DailyTaskSearch struct {
	Name     *string `json:"name" form:"name"`
	IsActive *int    `json:"isActive" form:"isActive"`
	request.PageInfo
}
