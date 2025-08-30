package miser

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/middleware"
)

type MiserCategoryRouter struct{}

func (s *MiserCategoryRouter) InitMiserCategoryRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	miserCategoryRouter := Router.Group("miserCategory").Use(middleware.OperationRecord())
	{
		miserCategoryRouter.POST("createMiserCategory", miserCategoryApi.CreateMiserCategory)
		miserCategoryRouter.DELETE("deleteMiserCategory", miserCategoryApi.DeleteMiserCategory)
		miserCategoryRouter.DELETE("deleteMiserCategoryByIds", miserCategoryApi.DeleteMiserCategoryByIds)
		miserCategoryRouter.PUT("updateMiserCategory", miserCategoryApi.UpdateMiserCategory)
	}

	miserCategoryRouterWithoutRecord := Router.Group("miserCategory")
	{
		miserCategoryRouterWithoutRecord.GET("findMiserCategory", miserCategoryApi.FindMiserCategory)
		miserCategoryRouterWithoutRecord.GET("getMiserCategoryList", miserCategoryApi.GetMiserCategoryList)
		miserCategoryRouterWithoutRecord.GET("listMiserCategoryList", miserCategoryApi.ListMiserCategoryList)
	}

	miserCategoryRouterWithoutAuth := PublicRouter.Group("miserCategory")
	{
		_ = miserCategoryRouterWithoutAuth
	}
}
