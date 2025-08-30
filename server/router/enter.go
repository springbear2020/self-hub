package router

import (
	"github.com/springbear2020/self-hub/server/router/example"
	"github.com/springbear2020/self-hub/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}
