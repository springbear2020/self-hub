package request

import (
	"github.com/springbear2020/self-hub/server/model/common/request"
)

type MiserCategorySearch struct {
	TransactionType *int    `json:"transactionType" form:"transactionType"`
	Name            *string `json:"name" form:"name"`
	request.PageInfo
}
