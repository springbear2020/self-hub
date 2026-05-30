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

type MineArchivesApi struct{}

func (mineArchivesApi *MineArchivesApi) CreateMineArchives(c *gin.Context) {
	var mineArchives mine.MineArchives
	err := c.ShouldBindJSON(&mineArchives)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = mineArchivesService.CreateMineArchives(uid, &mineArchives)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (mineArchivesApi *MineArchivesApi) DeleteMineArchives(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	err := mineArchivesService.DeleteMineArchives(uid, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (mineArchivesApi *MineArchivesApi) DeleteMineArchivesByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	uid := utils.GetUserID(c)
	err := mineArchivesService.DeleteMineArchivesByIds(uid, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("批量删除成功", c)
}

func (mineArchivesApi *MineArchivesApi) UpdateMineArchives(c *gin.Context) {
	var mineArchives mine.MineArchives
	err := c.ShouldBindJSON(&mineArchives)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = mineArchivesService.UpdateMineArchives(uid, mineArchives)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

func (mineArchivesApi *MineArchivesApi) FindMineArchives(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	remineArchives, err := mineArchivesService.GetMineArchives(uid, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(remineArchives, c)
}

func (mineArchivesApi *MineArchivesApi) GetMineArchivesList(c *gin.Context) {
	var pageInfo request.MineArchivesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.PageInfo.Check()

	uid := utils.GetUserID(c)
	list, total, err := mineArchivesService.GetMineArchivesInfoList(uid, pageInfo)
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
