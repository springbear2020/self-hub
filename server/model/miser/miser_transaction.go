package miser

import (
	"time"
)

type MiserTransaction struct {
	Id              *int       `json:"id" form:"id" gorm:"primarykey;comment:ID;column:id;size:10;"`
	UserId          uint       `json:"userId" form:"userId" gorm:"comment:用户ID;column:user_id;size:10;"`
	CategoryId      *int       `json:"categoryId" form:"categoryId" gorm:"comment:分类ID;column:category_id;size:10;" binding:"required"`
	TransactionType *int       `json:"transactionType" form:"transactionType" gorm:"comment:类型(1-收入,2-支出);column:transaction_type;" binding:"required"`
	Date            *time.Time `json:"date" form:"date" gorm:"comment:日期;column:date;" binding:"required"`
	Amount          *float64   `json:"amount" form:"amount" gorm:"comment:金额;column:amount;size:10;" binding:"required"`
	CreatedAt       *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`
	UpdatedAt       *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`
}

func (MiserTransaction) TableName() string {
	return "miser_transactions"
}
