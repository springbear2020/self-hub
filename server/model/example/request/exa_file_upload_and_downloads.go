package request

import (
	"github.com/springbear2020/self-hub/server/model/common/request"
)

type ExaAttachmentCategorySearch struct {
	ClassId int `json:"classId" form:"classId"`
	request.PageInfo
}
