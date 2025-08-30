package miser

import "github.com/springbear2020/self-hub/server/service"

type ApiGroup struct {
	MiserCategoryApi
	MiserTransactionApi
	MiserTransactionItemApi
	MiserStatApi
}

var (
	miserCategoryService        = service.ServiceGroupApp.MiserServiceGroup.MiserCategoryService
	miserTransactionService     = service.ServiceGroupApp.MiserServiceGroup.MiserTransactionService
	miserTransactionItemService = service.ServiceGroupApp.MiserServiceGroup.MiserTransactionItemService
	miserStatService            = service.ServiceGroupApp.MiserServiceGroup.MiserStatService
)
