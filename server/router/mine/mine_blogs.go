package mine

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/middleware"
)

type MineBlogsRouter struct{}

func (s *MineBlogsRouter) InitMineBlogsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mineBlogsRouter := Router.Group("mineBlogs").Use(middleware.OperationRecord())
	mineBlogsRouterWithoutRecord := Router.Group("mineBlogs")
	mineBlogsRouterWithoutAuth := PublicRouter.Group("mineBlogs")
	{
		mineBlogsRouter.POST("createMineBlogs", mineBlogsApi.CreateMineBlogs)
		mineBlogsRouter.DELETE("deleteMineBlogs", mineBlogsApi.DeleteMineBlogs)
		mineBlogsRouter.DELETE("deleteMineBlogsByIds", mineBlogsApi.DeleteMineBlogsByIds)
		mineBlogsRouter.PUT("updateMineBlogs", mineBlogsApi.UpdateMineBlogs)
	}
	{
		mineBlogsRouterWithoutRecord.GET("findMineBlogs", mineBlogsApi.FindMineBlogs)
		mineBlogsRouterWithoutRecord.GET("getMineBlogsList", mineBlogsApi.GetMineBlogsList)
		mineBlogsRouterWithoutRecord.GET("getMineBlogsStat", mineBlogsApi.GetMineBlogsStat)
	}
	{
		_ = mineBlogsRouterWithoutAuth
	}
}
