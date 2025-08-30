package dto

type MiserStatCategoryAmount struct {
	Category int     `json:"category" gorm:"category"`
	Amount   float64 `json:"amount" gorm:"amount"`
}

type MiserStatMonthAmount struct {
	Month   string  `json:"month" gorm:"month"`
	Income  float64 `json:"income" gorm:"income"`
	Expense float64 `json:"expense" gorm:"expense"`
	Balance float64 `json:"balance" gorm:"balance"`
}

type MiserStatYearAmount struct {
	Year    string  `json:"year" gorm:"year"`
	Income  float64 `json:"income" gorm:"income"`
	Expense float64 `json:"expense" gorm:"expense"`
	Balance float64 `json:"balance" gorm:"balance"`
}

type MiserStatTransactionAmount struct {
	CategoryId int     `json:"categoryId" gorm:"categoryId"`
	Name       string  `json:"name" gorm:"name"`
	Amount     float64 `json:"amount" gorm:"amount"`
}
