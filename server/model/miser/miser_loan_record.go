package miser

import (
	"time"
)

type MiserLoanRecord struct {
	Id          *int       `json:"id" form:"id" gorm:"primarykey;comment:ID;column:id;size:10;"`                                                //ID
	UserId      uint       `json:"userId" form:"userId" gorm:"comment:用户ID;column:user_id;size:10;"`                                            //UID
	Name        *string    `json:"name" form:"name" gorm:"comment:借还对象;column:name;size:64;" binding:"required"`                                //借还对象
	LendDate    *time.Time `json:"lendDate" form:"lendDate" gorm:"comment:借出日期;column:lend_date;" binding:"required"`                           //借出日期
	LendAmount  *float64   `json:"lendAmount" form:"lendAmount" gorm:"comment:借出金额;column:lend_amount;size:10;" binding:"required"`             //借出金额
	RepayDate   *time.Time `json:"repayDate" form:"repayDate" gorm:"comment:归还日期;column:repay_date;"`                                           //归还日期
	RepayAmount *float64   `json:"repayAmount" form:"repayAmount" gorm:"comment:归还金额;column:repay_amount;size:10;"`                             //归还金额
	Description *string    `json:"description" form:"description" gorm:"comment:借还说明;column:description;size:512;" binding:"required"`          //借还说明
	FundStatus  *int       `json:"fundStatus" form:"fundStatus" gorm:"comment:资金状态(1-待还款,2-部分还款,3-已结清);column:fund_status;" binding:"required"` //资金状态
	CreatedAt   *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`                                           //创建时间
	UpdatedAt   *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`                                           //更新时间
}

func (MiserLoanRecord) TableName() string {
	return "miser_loan_records"
}
