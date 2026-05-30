package mine

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/common/response"
	"github.com/springbear2020/self-hub/server/model/mine"
	"github.com/springbear2020/self-hub/server/model/mine/request"
	"github.com/springbear2020/self-hub/server/utils"
	"go.uber.org/zap"
)

type MineWebsitesApi struct{}

func (mineWebsitesApi *MineWebsitesApi) CreateMineWebsites(c *gin.Context) {
	var mineWebsites mine.MineWebsites
	err := c.ShouldBindJSON(&mineWebsites)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = mineWebsitesService.CreateMineWebsites(uid, &mineWebsites)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (mineWebsitesApi *MineWebsitesApi) DeleteMineWebsites(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	err := mineWebsitesService.DeleteMineWebsites(uid, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (mineWebsitesApi *MineWebsitesApi) DeleteMineWebsitesByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	uid := utils.GetUserID(c)
	err := mineWebsitesService.DeleteMineWebsitesByIds(uid, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("批量删除成功", c)
}

func (mineWebsitesApi *MineWebsitesApi) UpdateMineWebsites(c *gin.Context) {
	var mineWebsites mine.MineWebsites
	err := c.ShouldBindJSON(&mineWebsites)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = mineWebsitesService.UpdateMineWebsites(uid, mineWebsites)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

func (mineWebsitesApi *MineWebsitesApi) FindMineWebsites(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	remineWebsites, err := mineWebsitesService.GetMineWebsites(uid, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(remineWebsites, c)
}

func (mineWebsitesApi *MineWebsitesApi) GetMineWebsitesList(c *gin.Context) {
	var pageInfo request.MineWebsitesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.PageInfo.Check()

	uid := utils.GetUserID(c)
	list, total, err := mineWebsitesService.GetMineWebsitesInfoList(uid, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (mineWebsitesApi *MineWebsitesApi) ListMineWebsites(c *gin.Context) {
	uid := utils.GetUserID(c)
	list, err := mineWebsitesService.ListMineWebsites(uid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(list, c)
}
