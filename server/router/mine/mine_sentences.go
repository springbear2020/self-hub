package mine

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/middleware"
)

type MineSentencesRouter struct{}

func (s *MineSentencesRouter) InitMineSentencesRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mineSentencesRouter := Router.Group("mineSentences").Use(middleware.OperationRecord())
	mineSentencesRouterWithoutRecord := Router.Group("mineSentences")
	mineSentencesRouterWithoutAuth := PublicRouter.Group("mineSentences")
	{
		mineSentencesRouter.POST("createMineSentences", mineSentencesApi.CreateMineSentences)
		mineSentencesRouter.DELETE("deleteMineSentences", mineSentencesApi.DeleteMineSentences)
		mineSentencesRouter.DELETE("deleteMineSentencesByIds", mineSentencesApi.DeleteMineSentencesByIds)
		mineSentencesRouter.PUT("updateMineSentences", mineSentencesApi.UpdateMineSentences)
	}
	{
		mineSentencesRouterWithoutRecord.GET("findMineSentences", mineSentencesApi.FindMineSentences)
		mineSentencesRouterWithoutRecord.GET("getMineSentencesList", mineSentencesApi.GetMineSentencesList)
		mineSentencesRouterWithoutRecord.GET("getMineSentencesStat", mineSentencesApi.GetMineSentencesStat)
	}
	{
		_ = mineSentencesRouterWithoutAuth
	}
}
