package request

type MiserStat struct {
	StartMonth string `json:"startMonth" form:"startMonth" binding:"required"`
	EndMonth   string `json:"endMonth" form:"endMonth" binding:"required"`
}
