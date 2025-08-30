package example

import (
	api "github.com/springbear2020/self-hub/server/api/v1"
)

type RouterGroup struct {
	CustomerRouter
	FileUploadAndDownloadRouter
	AttachmentCategoryRouter
}

var (
	exaCustomerApi              = api.ApiGroupApp.ExampleApiGroup.CustomerApi
	exaFileUploadAndDownloadApi = api.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi
	attachmentCategoryApi       = api.ApiGroupApp.ExampleApiGroup.AttachmentCategoryApi
)
