package v1

import (
	"github.com/springbear2020/self-hub/server/api/v1/example"
	"github.com/springbear2020/self-hub/server/api/v1/miser"
	"github.com/springbear2020/self-hub/server/api/v1/system"
	"github.com/springbear2020/self-hub/server/api/v1/task"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	TaskApiGroup    task.ApiGroup
	MiserApiGroup   miser.ApiGroup
}
