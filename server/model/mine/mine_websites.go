package mine

import "time"

type MineWebsites struct {
	Id        *int       `json:"id" form:"id" gorm:"primarykey;comment:ID;column:id;size:10;"`
	UserId    uint       `json:"userId" form:"userId" gorm:"comment:UID;column:user_id;size:10;"`
	Name      *string    `json:"name" form:"name" gorm:"comment:网站简称;column:name;size:32;" binding:"required"`
	Title     *string    `json:"title" form:"title" gorm:"comment:网站标题;column:title;size:64;" binding:"required"`
	Href      *string    `json:"href" form:"href" gorm:"comment:链接;column:href;size:256;" binding:"required"`
	Icon      *string    `json:"icon" form:"icon" gorm:"comment:封面;column:icon;size:256;" binding:"required"`
	Category  *string    `json:"category" form:"category" gorm:"comment:分类;column:category;size:16;" binding:"required"`
	SortValue *int       `json:"sortValue" form:"sortValue" gorm:"comment:排序值;column:sort_value;size:10;" binding:"required"`
	CreatedAt *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`
	UpdatedAt *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`
}

func (MineWebsites) TableName() string {
	return "mine_websites"
}
