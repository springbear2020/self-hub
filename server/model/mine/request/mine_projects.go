package request

import (
	"github.com/springbear2020/self-hub/server/model/common/request"
)

type MineProjectsSearch struct {
	Name        *string `json:"name" form:"name"`
	Title       *string `json:"title" form:"title"`
	Description *string `json:"description" form:"description"`
	request.PageInfo
}
