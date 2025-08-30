package miser

import (
	"time"
)

type MiserCategory struct {
	Id              *int       `json:"id" form:"id" gorm:"primarykey;comment:ID;column:id;size:10;"`                                                    //ID
	UserId          uint       `json:"userId" form:"userId" gorm:"comment:用户ID;column:user_id;size:10;"`                                                //UID
	TransactionType *int       `json:"transactionType" form:"transactionType" gorm:"comment:类型(1-收入,2-支出);column:transaction_type;" binding:"required"` //交易类型
	Name            *string    `json:"name" form:"name" gorm:"comment:分类名称;column:name;size:64;" binding:"required"`                                    //分类名称
	Description     *string    `json:"description" form:"description" gorm:"comment:分类描述;column:description;size:256;" binding:"required"`              //分类描述
	Sort            *int       `json:"sort" form:"sort" gorm:"comment:排序值;column:sort;size:10;" binding:"required"`                                     //排序值
	CreatedAt       *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`                                               //创建时间
	UpdatedAt       *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`                                               //更新时间
}

func (MiserCategory) TableName() string {
	return "miser_categories"
}
