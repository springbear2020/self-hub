package miser

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/middleware"
)

type MiserRankingRecordRouter struct{}

func (s *MiserRankingRecordRouter) InitMiserRankingRecordRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	miserRankingRecordRouter := Router.Group("miserRankingRecord").Use(middleware.OperationRecord())
	{
		miserRankingRecordRouter.POST("createMiserRankingRecord", miserRankingRecordApi.CreateMiserRankingRecord)
		miserRankingRecordRouter.DELETE("deleteMiserRankingRecord", miserRankingRecordApi.DeleteMiserRankingRecord)
		miserRankingRecordRouter.DELETE("deleteMiserRankingRecordByIds", miserRankingRecordApi.DeleteMiserRankingRecordByIds)
		miserRankingRecordRouter.PUT("updateMiserRankingRecord", miserRankingRecordApi.UpdateMiserRankingRecord)
	}

	miserRankingRecordRouterWithoutRecord := Router.Group("miserRankingRecord")
	{
		miserRankingRecordRouterWithoutRecord.GET("findMiserRankingRecord", miserRankingRecordApi.FindMiserRankingRecord)
		miserRankingRecordRouterWithoutRecord.GET("getMiserRankingRecordList", miserRankingRecordApi.GetMiserRankingRecordList)
	}

	miserRankingRecordRouterWithoutAuth := PublicRouter.Group("miserRankingRecord")
	{
		_ = miserRankingRecordRouterWithoutAuth
	}
}
