package request

import (
	"github.com/springbear2020/self-hub/server/model/common/request"
	"time"
)

type MiserTransactionSearch struct {
	CategoryId      *int       `json:"categoryId" form:"categoryId"`
	TransactionType *int       `json:"transactionType" form:"transactionType"`
	Date            *time.Time `json:"date" form:"date"`
	request.PageInfo
}
