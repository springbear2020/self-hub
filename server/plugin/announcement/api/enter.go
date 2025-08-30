package api

import "github.com/springbear2020/self-hub/server/plugin/announcement/service"

var (
	Api         = new(api)
	serviceInfo = service.Service.Info
)

type api struct{ Info info }
