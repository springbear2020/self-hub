package dto

type MineResourcesDTO struct {
	Id          uint   `json:"id" gorm:"id"`
	Type        string `json:"type" gorm:"type"`
	Title       string `json:"title" gorm:"title"`
	Description string `json:"description" gorm:"description"`
	Link        string `json:"link" gorm:"link"`
}
