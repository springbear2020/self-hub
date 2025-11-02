package miser

import (
	"time"
)

type MiserRankingRecord struct {
	Id              *int       `json:"id" form:"id" gorm:"primarykey;comment:ID;column:id;size:10;"`                                                    //ID
	UserId          uint       `json:"userId" form:"userId" gorm:"comment:用户ID;column:user_id;size:10;"`                                                //UID
	TransactionType *int       `json:"transactionType" form:"transactionType" gorm:"comment:类型(1-收入,2-支出);column:transaction_type;" binding:"required"` //交易类型
	CategoryId      *int       `json:"categoryId" form:"categoryId" gorm:"comment:分类ID;column:category_id;size:10;" binding:"required"`                 //交易分类
	Name            *string    `json:"name" form:"name" gorm:"comment:名称;column:name;size:64;" binding:"required"`                                      //名称
	Amount          *float64   `json:"amount" form:"amount" gorm:"comment:金额;column:amount;size:10;" binding:"required"`                                //交易金额
	Date            *time.Time `json:"date" form:"date" gorm:"comment:交易日期;column:date;" binding:"required"`                                            //交易日期
	CreatedAt       *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`                                               //创建时间
	UpdatedAt       *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`                                               //更新时间
}

func (MiserRankingRecord) TableName() string {
	return "miser_ranking_records"
}
