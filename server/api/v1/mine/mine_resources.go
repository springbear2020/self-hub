package mine

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/common/response"
	"github.com/springbear2020/self-hub/server/utils"
	"go.uber.org/zap"
)

type MineResourcesApi struct{}

func (api *MineResourcesApi) SearchMineResources(c *gin.Context) {
	uid := utils.GetUserID(c)
	list, err := mineResourcesService.SearchMineResources(uid, c.Query("keyword"))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(list, c)
}
