package mine

import api "github.com/springbear2020/self-hub/server/api/v1"

type RouterGroup struct {
	MineBooksRouter
	MineBlogsRouter
	MineProjectsRouter
	MineWebsitesRouter
	MineArchivesRouter
	MineSentencesRouter
}

var (
	mineBooksApi     = api.ApiGroupApp.MineApiGroup.MineBooksApi
	mineBlogsApi     = api.ApiGroupApp.MineApiGroup.MineBlogsApi
	mineProjectsApi  = api.ApiGroupApp.MineApiGroup.MineProjectsApi
	mineWebsitesApi  = api.ApiGroupApp.MineApiGroup.MineWebsitesApi
	mineArchivesApi  = api.ApiGroupApp.MineApiGroup.MineArchivesApi
	mineSentencesApi = api.ApiGroupApp.MineApiGroup.MineSentencesApi
)
