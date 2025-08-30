package router

import (
	"github.com/springbear2020/self-hub/server/router/example"
	"github.com/springbear2020/self-hub/server/router/miser"
	"github.com/springbear2020/self-hub/server/router/system"
	"github.com/springbear2020/self-hub/server/router/task"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Task    task.RouterGroup
	Miser   miser.RouterGroup
}
