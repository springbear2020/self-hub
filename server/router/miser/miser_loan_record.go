package miser

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/middleware"
)

type MiserLoanRecordRouter struct{}

func (s *MiserLoanRecordRouter) InitMiserLoanRecordRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	miserLoanRecordRouter := Router.Group("miserLoanRecord").Use(middleware.OperationRecord())
	{
		miserLoanRecordRouter.POST("createMiserLoanRecord", miserLoanRecordApi.CreateMiserLoanRecord)
		miserLoanRecordRouter.DELETE("deleteMiserLoanRecord", miserLoanRecordApi.DeleteMiserLoanRecord)
		miserLoanRecordRouter.DELETE("deleteMiserLoanRecordByIds", miserLoanRecordApi.DeleteMiserLoanRecordByIds)
		miserLoanRecordRouter.PUT("updateMiserLoanRecord", miserLoanRecordApi.UpdateMiserLoanRecord)
	}

	miserLoanRecordRouterWithoutRecord := Router.Group("miserLoanRecord")
	{
		miserLoanRecordRouterWithoutRecord.GET("findMiserLoanRecord", miserLoanRecordApi.FindMiserLoanRecord)
		miserLoanRecordRouterWithoutRecord.GET("getMiserLoanRecordList", miserLoanRecordApi.GetMiserLoanRecordList)
		miserLoanRecordRouterWithoutRecord.GET("listMiserLoanNameList", miserLoanRecordApi.ListMiserLoanNameList)
		miserLoanRecordRouterWithoutRecord.GET("getMiserLoanStatData", miserLoanRecordApi.GetMiserLoanStatData)
	}

	miserLoanRecordRouterWithoutAuth := PublicRouter.Group("miserLoanRecord")
	{
		_ = miserLoanRecordRouterWithoutAuth
	}
}
