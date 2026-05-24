package opening

import api "github.com/springbear2020/self-hub/server/api/v1"

type RouterGroup struct {
	OpeningRouter
}

var (
	mineBooksApi     = api.ApiGroupApp.MineApiGroup.MineBooksApi
	mineBlogsApi     = api.ApiGroupApp.MineApiGroup.MineBlogsApi
	mineProjectsApi  = api.ApiGroupApp.MineApiGroup.MineProjectsApi
	mineResourcesApi = api.ApiGroupApp.MineApiGroup.MineResourcesApi
	mineWebsitesApi  = api.ApiGroupApp.MineApiGroup.MineWebsitesApi
	dictionaryApi    = api.ApiGroupApp.SystemApiGroup.DictionaryApi
)
