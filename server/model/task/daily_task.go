package task

import (
	"time"
)

type DailyTask struct {
	Id          *int       `json:"id" form:"id" gorm:"primarykey;comment:ID;column:id;size:10;"`
	UserId      uint       `json:"userId" form:"userId" gorm:"comment:用户 ID;column:user_id;size:10;"`
	Name        *string    `json:"name" form:"name" gorm:"comment:任务名称;column:name;size:64;" binding:"required"`
	Description *string    `json:"description" form:"description" gorm:"comment:任务描述;column:description;size:256;"`
	IsActive    *int       `json:"isActive" form:"isActive" gorm:"comment:是否启用(1-启用,0-禁用);column:is_active;" binding:"required"`
	Sort        *int       `json:"sort" form:"sort" gorm:"comment:排序值;column:sort;size:10;" binding:"required"`
	CreatedAt   *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`
	UpdatedAt   *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`
}

func (DailyTask) TableName() string {
	return "daily_tasks"
}
