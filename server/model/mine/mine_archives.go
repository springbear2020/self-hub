package mine

import (
	"time"
)

type MineArchives struct {
	Id          *int       `json:"id" form:"id" gorm:"primarykey;comment:ID;column:id;size:10;"`
	UserId      uint       `json:"userId" form:"userId" gorm:"comment:UID;column:user_id;size:10;"`
	Title       *string    `json:"title" form:"title" gorm:"comment:标题;column:title;size:32;" binding:"required"`
	Category    *string    `json:"category" form:"category" gorm:"comment:分类;column:category;size:32;" binding:"required"`
	Description *string    `json:"description" form:"description" gorm:"comment:描述;column:description;size:128;" binding:"required"`
	Link        *string    `json:"link" form:"link" gorm:"comment:链接;column:link;size:256;" binding:"required"`
	SortValue   *int       `json:"sortValue" form:"sortValue" gorm:"comment:排序值;column:sort_value;size:10;" binding:"required"`
	CreatedAt   *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`
	UpdatedAt   *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`
}

func (MineArchives) TableName() string {
	return "mine_archives"
}
