package miser

import (
	"fmt"
	"github.com/springbear2020/self-hub/server/constants"
	"time"
)

type MiserLoanRecord struct {
	Id                 *int       `json:"id" form:"id" gorm:"primarykey;comment:ID;column:id;size:10;"`
	UserId             uint       `json:"userId" form:"userId" gorm:"comment:用户ID;column:user_id;size:10;"`
	Name               *string    `json:"name" form:"name" gorm:"comment:借还对象;column:name;size:64;" binding:"required"`
	LendDate           *time.Time `json:"lendDate" form:"lendDate" gorm:"comment:借出日期;column:lend_date;" binding:"required"`
	LendAmount         *float64   `json:"lendAmount" form:"lendAmount" gorm:"comment:借出金额;column:lend_amount;size:10;" binding:"required"`
	RepayDate          *time.Time `json:"repayDate" form:"repayDate" gorm:"comment:归还日期;column:repay_date;" binding:"required"`
	RepayAmount        *float64   `json:"repayAmount" form:"repayAmount" gorm:"comment:归还金额;column:repay_amount;size:10;" binding:"required"`
	Description        *string    `json:"description" form:"description" gorm:"comment:借还说明;column:description;size:512;" binding:"required"`
	FundStatus         int        `json:"fundStatus" form:"fundStatus" gorm:"comment:资金状态(1-待还款,2-部分还款,3-已结清);column:fund_status;"`
	SettlementDuration int        `json:"settlementDuration" form:"settlementDuration" gorm:"comment:结清耗天;column:settlement_duration;size:10;"`
	CreatedAt          *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`
	UpdatedAt          *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`
}

func (*MiserLoanRecord) TableName() string {
	return "miser_loan_records"
}

func (record *MiserLoanRecord) CalculateLoanRecord() error {
	lend := int(*record.LendAmount * 100)   // 借款金额（分）
	repay := int(*record.RepayAmount * 100) // 还款金额（分）
	if repay > lend {
		return fmt.Errorf("repay amount cannot be greater than lend amount")
	}

	diff := lend - repay
	switch {
	case diff == 0:
		record.FundStatus = constants.FundStatusRepaid
		record.SettlementDuration = int(record.RepayDate.Sub(*record.LendDate).Hours() / 24)
	case diff == lend:
		record.FundStatus = constants.FundStatusLoaning
	default:
		record.FundStatus = constants.FundStatusRepaying
	}

	return nil
}
