package request

import (
	"github.com/springbear2020/self-hub/server/model/common/request"
)

type MineBlogsSearch struct {
	Title      *string `json:"title" form:"title"`
	CategoryId *int    `json:"categoryId" form:"categoryId"`
	TagId      *int    `json:"tagId" form:"tagId"`
	request.PageInfo
}
