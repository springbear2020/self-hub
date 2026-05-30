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

type MineSentencesApi struct{}

func (mineSentencesApi *MineSentencesApi) CreateMineSentences(c *gin.Context) {
	var mineSentences mine.MineSentences
	err := c.ShouldBindJSON(&mineSentences)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = mineSentencesService.CreateMineSentences(uid, &mineSentences)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (mineSentencesApi *MineSentencesApi) DeleteMineSentences(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	err := mineSentencesService.DeleteMineSentences(uid, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (mineSentencesApi *MineSentencesApi) DeleteMineSentencesByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	uid := utils.GetUserID(c)
	err := mineSentencesService.DeleteMineSentencesByIds(uid, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("批量删除成功", c)
}

func (mineSentencesApi *MineSentencesApi) UpdateMineSentences(c *gin.Context) {
	var mineSentences mine.MineSentences
	err := c.ShouldBindJSON(&mineSentences)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = mineSentencesService.UpdateMineSentences(uid, mineSentences)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

func (mineSentencesApi *MineSentencesApi) FindMineSentences(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	remineSentences, err := mineSentencesService.GetMineSentences(uid, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(remineSentences, c)
}

func (mineSentencesApi *MineSentencesApi) GetMineSentencesList(c *gin.Context) {
	var pageInfo request.MineSentencesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.PageInfo.Check()

	uid := utils.GetUserID(c)
	list, total, err := mineSentencesService.GetMineSentencesInfoList(uid, pageInfo)
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

func (mineSentencesApi *MineSentencesApi) GetMineSentencesStat(c *gin.Context) {
	uid := utils.GetUserID(c)
	list, err := mineSentencesService.GetMineSentencesStat(uid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(list, c)
}
