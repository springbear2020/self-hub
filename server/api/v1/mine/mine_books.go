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

type MineBooksApi struct{}

func (mineBooksApi *MineBooksApi) CreateMineBooks(c *gin.Context) {
	var mineBooks mine.MineBooks
	err := c.ShouldBindJSON(&mineBooks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = mineBooksService.CreateMineBooks(uid, &mineBooks)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (mineBooksApi *MineBooksApi) DeleteMineBooks(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	err := mineBooksService.DeleteMineBooks(uid, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (mineBooksApi *MineBooksApi) DeleteMineBooksByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	uid := utils.GetUserID(c)
	err := mineBooksService.DeleteMineBooksByIds(uid, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("批量删除成功", c)
}

func (mineBooksApi *MineBooksApi) UpdateMineBooks(c *gin.Context) {
	var mineBooks mine.MineBooks
	err := c.ShouldBindJSON(&mineBooks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = mineBooksService.UpdateMineBooks(uid, mineBooks)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

func (mineBooksApi *MineBooksApi) FindMineBooks(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	remineBooks, err := mineBooksService.GetMineBooks(uid, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(remineBooks, c)
}

func (mineBooksApi *MineBooksApi) GetMineBooksList(c *gin.Context) {
	var pageInfo request.MineBooksSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.PageInfo.Check()

	uid := utils.GetUserID(c)
	list, total, err := mineBooksService.GetMineBooksInfoList(uid, pageInfo)
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
