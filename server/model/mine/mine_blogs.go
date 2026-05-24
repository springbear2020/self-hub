package mine

import (
	"time"
)

type MineBlogs struct {
	Id         *int       `json:"id" form:"id" gorm:"primarykey;comment:ID;column:id;size:10;"`
	UserId     uint       `json:"userId" form:"userId" gorm:"comment:UID;column:user_id;size:10;"`
	Title      *string    `json:"title" form:"title" gorm:"comment:标题;column:title;size:64;"`
	Link       *string    `json:"link" form:"link" gorm:"comment:链接;column:link;size:256;"`
	Cover      *string    `json:"cover" form:"cover" gorm:"comment:封面;column:cover;size:256;"`
	PostTime   *time.Time `json:"postTime" form:"postTime" gorm:"comment:发布时间;column:post_time;"`
	CategoryId *int       `json:"categoryId" form:"categoryId" gorm:"comment:分类;column:category_id;size:10;"`
	TagId      *int       `json:"tagId" form:"tagId" gorm:"comment:标签;column:tag_id;size:10;"`
	SortValue  *int       `json:"sortValue" form:"sortValue" gorm:"comment:排序值;column:sort_value;size:10;"`
	CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`
	UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`
}

func (MineBlogs) TableName() string {
	return "mine_blogs"
}
