package mine

import "github.com/springbear2020/self-hub/server/service/system"

type ServiceGroup struct {
	MineBooksService
	MineBlogsService
	MineProjectsService
	MineResourcesService
	MineWebsitesService
}

var (
	mineBooksService    = MineBooksService{}
	mineBlogsService    = MineBlogsService{}
	mineProjectsService = MineProjectsService{}
	dictionaryService   = system.DictionaryService{}
)
