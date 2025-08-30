package request

import (
	"github.com/springbear2020/self-hub/server/model/common/request"
	"github.com/springbear2020/self-hub/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
