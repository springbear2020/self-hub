package miser

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/middleware"
)

type MiserTransactionRouter struct{}

func (s *MiserTransactionRouter) InitMiserTransactionRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	miserTransactionRouter := Router.Group("miserTransaction").Use(middleware.OperationRecord())
	{
		miserTransactionRouter.POST("createMiserTransaction", miserTransactionApi.CreateMiserTransaction)
		miserTransactionRouter.DELETE("deleteMiserTransaction", miserTransactionApi.DeleteMiserTransaction)
		miserTransactionRouter.DELETE("deleteMiserTransactionByIds", miserTransactionApi.DeleteMiserTransactionByIds)
		miserTransactionRouter.PUT("updateMiserTransaction", miserTransactionApi.UpdateMiserTransaction)
	}

	miserTransactionRouterWithoutRecord := Router.Group("miserTransaction")
	{
		miserTransactionRouterWithoutRecord.GET("findMiserTransaction", miserTransactionApi.FindMiserTransaction)
		miserTransactionRouterWithoutRecord.GET("getMiserTransactionList", miserTransactionApi.GetMiserTransactionList)
	}

	miserTransactionRouterWithoutAuth := PublicRouter.Group("miserTransaction")
	{
		_ = miserTransactionRouterWithoutAuth
	}
}
