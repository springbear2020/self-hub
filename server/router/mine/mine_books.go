package mine

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/middleware"
)

type MineBooksRouter struct{}

func (s *MineBooksRouter) InitMineBooksRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mineBooksRouter := Router.Group("mineBooks").Use(middleware.OperationRecord())
	mineBooksRouterWithoutRecord := Router.Group("mineBooks")
	mineBooksRouterWithoutAuth := PublicRouter.Group("mineBooks")
	{
		mineBooksRouter.POST("createMineBooks", mineBooksApi.CreateMineBooks)
		mineBooksRouter.DELETE("deleteMineBooks", mineBooksApi.DeleteMineBooks)
		mineBooksRouter.DELETE("deleteMineBooksByIds", mineBooksApi.DeleteMineBooksByIds)
		mineBooksRouter.PUT("updateMineBooks", mineBooksApi.UpdateMineBooks)
	}
	{
		mineBooksRouterWithoutRecord.GET("findMineBooks", mineBooksApi.FindMineBooks)
		mineBooksRouterWithoutRecord.GET("getMineBooksList", mineBooksApi.GetMineBooksList)
	}
	{
		_ = mineBooksRouterWithoutAuth
	}
}
