package request

import (
	"github.com/springbear2020/self-hub/server/model/common/request"
	"github.com/springbear2020/self-hub/server/model/system"
	"time"
)

type SysExportTemplateSearch struct {
	system.SysExportTemplate
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
