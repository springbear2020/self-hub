package miser

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/middleware"
)

type MiserTransactionItemRouter struct{}

func (s *MiserTransactionItemRouter) InitMiserTransactionItemRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	miserTransactionItemRouter := Router.Group("miserTransactionItem").Use(middleware.OperationRecord())
	{
		miserTransactionItemRouter.POST("createMiserTransactionItem", miserTransactionItemApi.CreateMiserTransactionItem)
		miserTransactionItemRouter.DELETE("deleteMiserTransactionItem", miserTransactionItemApi.DeleteMiserTransactionItem)
		miserTransactionItemRouter.DELETE("deleteMiserTransactionItemByIds", miserTransactionItemApi.DeleteMiserTransactionItemByIds)
		miserTransactionItemRouter.PUT("updateMiserTransactionItem", miserTransactionItemApi.UpdateMiserTransactionItem)
	}

	miserTransactionItemRouterWithoutRecord := Router.Group("miserTransactionItem")
	{
		miserTransactionItemRouterWithoutRecord.GET("findMiserTransactionItem", miserTransactionItemApi.FindMiserTransactionItem)
		miserTransactionItemRouterWithoutRecord.GET("getMiserTransactionItemList", miserTransactionItemApi.GetMiserTransactionItemList)
		miserTransactionItemRouterWithoutRecord.GET("listItemDistinctNames", miserTransactionItemApi.ListItemDistinctNames)
	}

	miserTransactionItemRouterWithoutAuth := PublicRouter.Group("miserTransactionItem")
	{
		_ = miserTransactionItemRouterWithoutAuth
	}
}
