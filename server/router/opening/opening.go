package opening

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/middleware"
)

type OpeningRouter struct{}

func (s *OpeningRouter) InitOpeningRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	withoutAuth := PublicRouter.Group("open").Use(middleware.OpenCheck())
	{
		withoutAuth.GET("mine/books", mineBooksApi.GetMineBooksList)            // books page
		withoutAuth.GET("mine/blogs", mineBlogsApi.GetMineBlogsList)            // blogs page
		withoutAuth.GET("mine/projects", mineProjectsApi.GetMineProjectsList)   // projects page
		withoutAuth.GET("mine/resources", mineResourcesApi.SearchMineResources) // search: books + blogs + projects
		withoutAuth.GET("mine/websites", mineWebsitesApi.ListMineWebsites)      // websites category list
		withoutAuth.GET("dict/list", dictionaryApi.ListArray)                   // dictionary list
		withoutAuth.GET("dict/map", dictionaryApi.GetMap)                       // dictionary map
	}
}
