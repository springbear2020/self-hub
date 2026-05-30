package example

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/middleware"
)

type DashboardRouter struct{}

func (e *CustomerRouter) InitDashboardRouter(Router *gin.RouterGroup) {
	dashboardRouter := Router.Group("dashboard").Use(middleware.OperationRecord())
	dashboardRouterWithoutRecord := Router.Group("dashboard")
	{
		_ = dashboardRouter
	}
	{
		dashboardRouterWithoutRecord.GET("stat", dashboardApi.GetDashboardStatData)
	}
}
