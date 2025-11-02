package request

type MiserStat struct {
	StartMonth string `json:"startMonth" form:"startMonth" binding:"required"`
	EndMonth   string `json:"endMonth" form:"endMonth" binding:"required"`
}

type MiserCategoryStat struct {
	MiserStat
	CategoryId int `json:"categoryId" form:"categoryId"`
}

type MiserTransactionStat struct {
	MiserStat
	TransactionType int `json:"transactionType" form:"transactionType"`
}

type MiserCategoryItemStat struct {
	MiserStat
	CategoryId int    `json:"categoryId" form:"categoryId"`
	Name       string `json:"name" form:"name"`
}
