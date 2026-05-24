package mine

import (
	"time"
)

type MineProjects struct {
	Id          *int       `json:"id" form:"id" gorm:"primarykey;comment:ID;column:id;size:10;"`
	UserId      uint       `json:"userId" form:"userId" gorm:"comment:UID;column:user_id;size:10;"`
	Name        *string    `json:"name" form:"name" gorm:"comment:项目英文名;column:name;size:32;"`
	Title       *string    `json:"title" form:"title" gorm:"comment:项目中文名;column:title;size:32;"`
	Technology  *string    `json:"technology" form:"technology" gorm:"comment:核心技术;column:technology;size:64;"`
	Description *string    `json:"description" form:"description" gorm:"comment:项目描述;column:description;size:256;"`
	Cover       *string    `json:"cover" form:"cover" gorm:"comment:封面图片;column:cover;size:256;"`
	Link        *string    `json:"link" form:"link" gorm:"comment:项目链接;column:link;size:256;"`
	StartDate   *time.Time `json:"startDate" form:"startDate" gorm:"comment:开始时间;column:start_date;"`
	FinishDate  *time.Time `json:"finishDate" form:"finishDate" gorm:"comment:完成时间;column:finish_date;"`
	SortValue   *int       `json:"sortValue" form:"sortValue" gorm:"comment:排序值;column:sort_value;size:10;"`
	CreatedAt   *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`
	UpdatedAt   *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`
}

func (MineProjects) TableName() string {
	return "mine_projects"
}
