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

type MineProjectsApi struct{}

func (mineProjectsApi *MineProjectsApi) CreateMineProjects(c *gin.Context) {
	var mineProjects mine.MineProjects
	err := c.ShouldBindJSON(&mineProjects)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = mineProjectsService.CreateMineProjects(uid, &mineProjects)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (mineProjectsApi *MineProjectsApi) DeleteMineProjects(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	err := mineProjectsService.DeleteMineProjects(uid, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (mineProjectsApi *MineProjectsApi) DeleteMineProjectsByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	uid := utils.GetUserID(c)
	err := mineProjectsService.DeleteMineProjectsByIds(uid, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("批量删除成功", c)
}

func (mineProjectsApi *MineProjectsApi) UpdateMineProjects(c *gin.Context) {
	var mineProjects mine.MineProjects
	err := c.ShouldBindJSON(&mineProjects)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = mineProjectsService.UpdateMineProjects(uid, mineProjects)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

func (mineProjectsApi *MineProjectsApi) FindMineProjects(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	remineProjects, err := mineProjectsService.GetMineProjects(uid, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(remineProjects, c)
}

func (mineProjectsApi *MineProjectsApi) GetMineProjectsList(c *gin.Context) {
	var pageInfo request.MineProjectsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.PageInfo.Check()

	uid := utils.GetUserID(c)
	list, total, err := mineProjectsService.GetMineProjectsInfoList(uid, pageInfo)
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
