package request

import (
	"github.com/springbear2020/self-hub/server/model/common/request"
)

type MineSentencesSearch struct {
	Title      *string `json:"title" form:"title"`
	Author     *string `json:"author" form:"author"`
	Dynasty    *string `json:"dynasty" form:"dynasty"`
	Content    *string `json:"content" form:"content"`
	CategoryId *int    `json:"categoryId" form:"categoryId"`
	request.PageInfo
}
