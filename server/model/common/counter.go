package common

type CountDTO struct {
	Label int `json:"label" gorm:"column:label"`
	Total int `json:"total" gorm:"column:total"`
}
