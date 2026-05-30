package mine

import "github.com/springbear2020/self-hub/server/service"

type ApiGroup struct {
	MineBooksApi
	MineBlogsApi
	MineProjectsApi
	MineResourcesApi
	MineWebsitesApi
	MineArchivesApi
	MineSentencesApi
}

var (
	mineBooksService     = service.ServiceGroupApp.MineServiceGroup.MineBooksService
	mineBlogsService     = service.ServiceGroupApp.MineServiceGroup.MineBlogsService
	mineProjectsService  = service.ServiceGroupApp.MineServiceGroup.MineProjectsService
	mineResourcesService = service.ServiceGroupApp.MineServiceGroup.MineResourcesService
	mineWebsitesService  = service.ServiceGroupApp.MineServiceGroup.MineWebsitesService
	mineArchivesService  = service.ServiceGroupApp.MineServiceGroup.MineArchivesService
	mineSentencesService = service.ServiceGroupApp.MineServiceGroup.MineSentencesService
)
