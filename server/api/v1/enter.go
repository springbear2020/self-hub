package v1

import (
	"github.com/springbear2020/self-hub/server/api/v1/example"
	"github.com/springbear2020/self-hub/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}
