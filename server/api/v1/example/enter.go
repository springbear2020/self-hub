package example

import "github.com/springbear2020/self-hub/server/service"

type ApiGroup struct {
	CustomerApi
	FileUploadAndDownloadApi
	AttachmentCategoryApi
}

var (
	customerService              = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
	attachmentCategoryService    = service.ServiceGroupApp.ExampleServiceGroup.AttachmentCategoryService
)
