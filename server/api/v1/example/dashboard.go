package example

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/common/response"
	"github.com/springbear2020/self-hub/server/utils"
	"go.uber.org/zap"
)

type DashboardApi struct{}

func (api *DashboardApi) GetDashboardStatData(c *gin.Context) {
	uid := utils.GetUserID(c)
	data, err := dashboardService.GetDashboardStatData(uid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}
