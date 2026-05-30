package mine

import (
	"time"
)

type MineSentences struct {
	Id         *int       `json:"id" form:"id" gorm:"primarykey;comment:ID;column:id;size:10;"`
	UserId     uint       `json:"userId" form:"userId" gorm:"comment:UID;column:user_id;size:10;"`
	Title      *string    `json:"title" form:"title" gorm:"comment:标题;column:title;size:64;" binding:"required"`
	Author     *string    `json:"author" form:"author" gorm:"comment:作者;column:author;size:32;" binding:"required"`
	Dynasty    *string    `json:"dynasty" form:"dynasty" gorm:"comment:朝代;column:dynasty;size:16;" binding:"required"`
	Content    *string    `json:"content" form:"content" gorm:"comment:内容;column:content;size:1024;" binding:"required"`
	Link       *string    `json:"link" form:"link" gorm:"comment:链接;column:link;size:256;" binding:"required"`
	CategoryId *int       `json:"categoryId" form:"categoryId" gorm:"comment:分类;column:category_id;size:10;" binding:"required"`
	SortValue  *int       `json:"sortValue" form:"sortValue" gorm:"comment:排序值;column:sort_value;size:10;" binding:"required"`
	CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`
	UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`
}

func (MineSentences) TableName() string {
	return "mine_sentences"
}
