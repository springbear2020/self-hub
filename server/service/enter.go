package service

import (
	"github.com/springbear2020/self-hub/server/service/example"
	"github.com/springbear2020/self-hub/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}
