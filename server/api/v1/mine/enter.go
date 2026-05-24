package mine

import "github.com/springbear2020/self-hub/server/service"

type ApiGroup struct {
	MineBooksApi
	MineBlogsApi
	MineProjectsApi
	MineResourcesApi
	MineWebsitesApi
}

var (
	mineBooksService     = service.ServiceGroupApp.MineServiceGroup.MineBooksService
	mineBlogsService     = service.ServiceGroupApp.MineServiceGroup.MineBlogsService
	mineProjectsService  = service.ServiceGroupApp.MineServiceGroup.MineProjectsService
	mineResourcesService = service.ServiceGroupApp.MineServiceGroup.MineResourcesService
	mineWebsitesService  = service.ServiceGroupApp.MineServiceGroup.MineWebsitesService
)
