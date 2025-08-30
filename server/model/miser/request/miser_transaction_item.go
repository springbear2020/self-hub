package request

import (
	"github.com/springbear2020/self-hub/server/model/common/request"
)

type MiserTransactionItemSearch struct {
	TransactionId *int    `json:"transactionId" form:"transactionId"`
	CategoryId    *int    `json:"categoryId" form:"categoryId"`
	Name          *string `json:"name" form:"name"`
	request.PageInfo
}
