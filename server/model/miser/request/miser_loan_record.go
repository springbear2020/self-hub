package request

import (
	"github.com/springbear2020/self-hub/server/model/common/request"
	"time"
)

type MiserLoanRecordSearch struct {
	Name       *string    `json:"name" form:"name"`
	LendDate   *time.Time `json:"lendDate" form:"lendDate"`
	RepayDate  *time.Time `json:"repayDate" form:"repayDate"`
	FundStatus *int       `json:"fundStatus" form:"fundStatus"`
	request.PageInfo
}
