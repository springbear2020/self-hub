package request

import (
	"github.com/springbear2020/self-hub/server/model/common/request"
)

type MineArchivesSearch struct {
	Title       *string `json:"title" form:"title"`
	Category    *string `json:"category" form:"category"`
	Description *string `json:"description" form:"description"`
	request.PageInfo
}
