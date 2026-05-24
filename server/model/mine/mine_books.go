package mine

import (
	"time"
)

type MineBooks struct {
	Id              *int       `json:"id" form:"id" gorm:"primarykey;comment:ID;column:id;size:10;"`
	UserId          uint       `json:"userId" form:"userId" gorm:"comment:UID;column:user_id;size:10;"`
	Title           *string    `json:"title" form:"title" gorm:"comment:书名;column:title;size:32;"`
	Author          *string    `json:"author" form:"author" gorm:"comment:作者;column:author;size:32;"`
	Cover           *string    `json:"cover" form:"cover" gorm:"comment:封面;column:cover;size:256;"`
	Link            *string    `json:"link" form:"link" gorm:"comment:豆瓣链接;column:link;size:256;"`
	FirstCompletion *time.Time `json:"firstCompletion" form:"firstCompletion" gorm:"comment:首次读完;column:first_completion;"`
	ReadMinute      *int       `json:"readMinute" form:"readMinute" gorm:"comment:阅读时长;column:read_minute;size:10;"`
	IconicQuote     *string    `json:"iconicQuote" form:"iconicQuote" gorm:"comment:经典名句;column:iconic_quote;size:256;"`
	SortValue       *int       `json:"sortValue" form:"sortValue" gorm:"comment:排序值;column:sort_value;size:10;"`
	CreatedAt       *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`
	UpdatedAt       *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`
}

func (MineBooks) TableName() string {
	return "mine_books"
}
