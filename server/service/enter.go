package service

import (
	"github.com/springbear2020/self-hub/server/service/example"
	"github.com/springbear2020/self-hub/server/service/miser"
	"github.com/springbear2020/self-hub/server/service/system"
	"github.com/springbear2020/self-hub/server/service/task"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	TaskServiceGroup    task.ServiceGroup
	MiserServiceGroup   miser.ServiceGroup
}
