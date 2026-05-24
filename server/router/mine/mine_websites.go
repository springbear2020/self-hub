package mine

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/middleware"
)

type MineWebsitesRouter struct{}

func (s *MineWebsitesRouter) InitMineWebsitesRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mineWebsitesRouter := Router.Group("mineWebsites").Use(middleware.OperationRecord())
	mineWebsitesRouterWithoutRecord := Router.Group("mineWebsites")
	mineWebsitesRouterWithoutAuth := PublicRouter.Group("mineWebsites")
	{
		mineWebsitesRouter.POST("createMineWebsites", mineWebsitesApi.CreateMineWebsites)
		mineWebsitesRouter.DELETE("deleteMineWebsites", mineWebsitesApi.DeleteMineWebsites)
		mineWebsitesRouter.DELETE("deleteMineWebsitesByIds", mineWebsitesApi.DeleteMineWebsitesByIds)
		mineWebsitesRouter.PUT("updateMineWebsites", mineWebsitesApi.UpdateMineWebsites)
	}
	{
		mineWebsitesRouterWithoutRecord.GET("findMineWebsites", mineWebsitesApi.FindMineWebsites)
		mineWebsitesRouterWithoutRecord.GET("getMineWebsitesList", mineWebsitesApi.GetMineWebsitesList)
	}
	{
		_ = mineWebsitesRouterWithoutAuth
	}
}
