package initialize

import (
	_ "github.com/springbear2020/self-hub/server/source/example"
	_ "github.com/springbear2020/self-hub/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
