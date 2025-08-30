package miser

import "github.com/gin-gonic/gin"

type MiserStatRouter struct {
}

func (s *MiserCategoryRouter) InitMiserStatRouter(privateRouter *gin.RouterGroup, publicRouter *gin.RouterGroup) {
	withoutRecord := privateRouter.Group("miserStat")
	{
		withoutRecord.GET("getCardStat", miserStatApi.GetCardStat)
		withoutRecord.GET("getPieStat", miserStatApi.GetPieStat)
		withoutRecord.GET("getLineStat", miserStatApi.GetLineStat)
		withoutRecord.GET("getMonthStat", miserStatApi.GetMonthStat)
		withoutRecord.GET("getYearStat", miserStatApi.GetYearStat)
		withoutRecord.GET("getItemStat", miserStatApi.GetItemStat)
	}

	withoutAuth := publicRouter.Group("miserStat")
	{
		_ = withoutAuth
	}
}
