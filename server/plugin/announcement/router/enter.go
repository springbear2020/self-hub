package router

import "github.com/springbear2020/self-hub/server/plugin/announcement/api"

var (
	Router  = new(router)
	apiInfo = api.Api.Info
)

type router struct{ Info info }
