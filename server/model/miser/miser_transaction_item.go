package miser

import (
	"time"
)

type MiserTransactionItem struct {
	Id            *int       `json:"id" form:"id" gorm:"primarykey;comment:ID;column:id;size:10;"`
	UserId        uint       `json:"userId" form:"userId" gorm:"comment:用户 ID;column:user_id;size:10;"`
	TransactionId *int       `json:"transactionId" form:"transactionId" gorm:"comment:交易 ID;column:transaction_id;size:10;" binding:"required"`
	CategoryId    *int       `json:"categoryId" form:"categoryId" gorm:"comment:分类;column:category_id;size:10;" binding:"required"`
	Name          *string    `json:"name" form:"name" gorm:"comment:名称;column:name;size:64;" binding:"required"`
	Amount        *float64   `json:"amount" form:"amount" gorm:"comment:金额;column:amount;size:10;" binding:"required"`
	Date          *time.Time `json:"date" form:"date" gorm:"comment:交易日期;column:date;" binding:"required"`
	CreatedAt     *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`
	UpdatedAt     *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`
}

func (MiserTransactionItem) TableName() string {
	return "miser_transaction_items"
}
