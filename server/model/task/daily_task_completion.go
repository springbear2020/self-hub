package task

import (
	"time"
)

type DailyTaskCompletion struct {
	Id         *int       `json:"id" form:"id" gorm:"primarykey;comment:ID;column:id;size:10;"`
	UserId     uint       `json:"userId" form:"userId" gorm:"comment:用户 ID;column:user_id;size:10;"`
	TaskId     *int       `json:"taskId" form:"taskId" gorm:"comment:任务 ID;column:task_id;size:10;" binding:"required"`
	FinishDate *time.Time `json:"finishDate" form:"finishDate" gorm:"comment:完成时间;column:finish_date;" binding:"required"`
	CountValue *int       `json:"countValue" form:"countValue" gorm:"comment:完成计数值;column:count_value;size:5;" binding:"required"`
	Remark     *string    `json:"remark" form:"remark" gorm:"comment:完成说明;column:remark;size:256;"`
	CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`
	UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`
}

func (DailyTaskCompletion) TableName() string {
	return "daily_task_completions"
}
