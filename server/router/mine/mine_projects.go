package mine

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/middleware"
)

type MineProjectsRouter struct{}

func (s *MineProjectsRouter) InitMineProjectsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mineProjectsRouter := Router.Group("mineProjects").Use(middleware.OperationRecord())
	mineProjectsRouterWithoutRecord := Router.Group("mineProjects")
	mineProjectsRouterWithoutAuth := PublicRouter.Group("mineProjects")
	{
		mineProjectsRouter.POST("createMineProjects", mineProjectsApi.CreateMineProjects)
		mineProjectsRouter.DELETE("deleteMineProjects", mineProjectsApi.DeleteMineProjects)
		mineProjectsRouter.DELETE("deleteMineProjectsByIds", mineProjectsApi.DeleteMineProjectsByIds)
		mineProjectsRouter.PUT("updateMineProjects", mineProjectsApi.UpdateMineProjects)
	}
	{
		mineProjectsRouterWithoutRecord.GET("findMineProjects", mineProjectsApi.FindMineProjects)
		mineProjectsRouterWithoutRecord.GET("getMineProjectsList", mineProjectsApi.GetMineProjectsList)
	}
	{
		_ = mineProjectsRouterWithoutAuth
	}
}
