package miser

import api "github.com/springbear2020/self-hub/server/api/v1"

type RouterGroup struct {
	MiserCategoryRouter
	MiserTransactionRouter
	MiserTransactionItemRouter
	MiserStatRouter
	MiserLoanRecordRouter
	MiserRankingRecordRouter
}

var (
	miserCategoryApi        = api.ApiGroupApp.MiserApiGroup.MiserCategoryApi
	miserTransactionApi     = api.ApiGroupApp.MiserApiGroup.MiserTransactionApi
	miserTransactionItemApi = api.ApiGroupApp.MiserApiGroup.MiserTransactionItemApi
	miserStatApi            = api.ApiGroupApp.MiserApiGroup.MiserStatApi
	miserLoanRecordApi      = api.ApiGroupApp.MiserApiGroup.MiserLoanRecordApi
	miserRankingRecordApi   = api.ApiGroupApp.MiserApiGroup.MiserRankingRecordApi
)
