package request

import (
	"github.com/springbear2020/self-hub/server/model/common/request"
)

type MineWebsitesSearch struct {
	Name     *string `json:"name" form:"name"`
	Title    *string `json:"title" form:"title"`
	Category *string `json:"category" form:"category"`
	request.PageInfo
}
