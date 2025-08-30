package request

import (
	"github.com/springbear2020/self-hub/server/model/common/request"
	"github.com/springbear2020/self-hub/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
