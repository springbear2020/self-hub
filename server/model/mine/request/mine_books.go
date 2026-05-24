package request

import (
	"github.com/springbear2020/self-hub/server/model/common/request"
)

type MineBooksSearch struct {
	Title       *string `json:"title" form:"title"`
	Author      *string `json:"author" form:"author"`
	IconicQuote *string `json:"iconicQuote" form:"iconicQuote"`
	request.PageInfo
}
