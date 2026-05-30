package mine

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/middleware"
)

type MineArchivesRouter struct{}

func (s *MineArchivesRouter) InitMineArchivesRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mineArchivesRouter := Router.Group("mineArchives").Use(middleware.OperationRecord())
	mineArchivesRouterWithoutRecord := Router.Group("mineArchives")
	mineArchivesRouterWithoutAuth := PublicRouter.Group("mineArchives")
	{
		mineArchivesRouter.POST("createMineArchives", mineArchivesApi.CreateMineArchives)
		mineArchivesRouter.DELETE("deleteMineArchives", mineArchivesApi.DeleteMineArchives)
		mineArchivesRouter.DELETE("deleteMineArchivesByIds", mineArchivesApi.DeleteMineArchivesByIds)
		mineArchivesRouter.PUT("updateMineArchives", mineArchivesApi.UpdateMineArchives)
	}
	{
		mineArchivesRouterWithoutRecord.GET("findMineArchives", mineArchivesApi.FindMineArchives)
		mineArchivesRouterWithoutRecord.GET("getMineArchivesList", mineArchivesApi.GetMineArchivesList)
	}
	{
		_ = mineArchivesRouterWithoutAuth
	}
}
