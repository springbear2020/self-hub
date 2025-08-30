package miser

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/common/response"
	"github.com/springbear2020/self-hub/server/model/miser"
	"github.com/springbear2020/self-hub/server/model/miser/request"
	"github.com/springbear2020/self-hub/server/utils"
	"go.uber.org/zap"
)

type MiserCategoryApi struct{}

func (miserCategoryApi *MiserCategoryApi) CreateMiserCategory(c *gin.Context) {
	var miserCategory miser.MiserCategory
	err := c.ShouldBindJSON(&miserCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = miserCategoryService.CreateMiserCategory(uid, &miserCategory)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (miserCategoryApi *MiserCategoryApi) DeleteMiserCategory(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	err := miserCategoryService.DeleteMiserCategory(uid, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (miserCategoryApi *MiserCategoryApi) DeleteMiserCategoryByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	uid := utils.GetUserID(c)
	err := miserCategoryService.DeleteMiserCategoryByIds(uid, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("批量删除成功", c)
}

func (miserCategoryApi *MiserCategoryApi) UpdateMiserCategory(c *gin.Context) {
	var miserCategory miser.MiserCategory
	err := c.ShouldBindJSON(&miserCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	err = miserCategoryService.UpdateMiserCategory(uid, miserCategory)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

func (miserCategoryApi *MiserCategoryApi) FindMiserCategory(c *gin.Context) {
	id := c.Query("id")
	uid := utils.GetUserID(c)
	miserCategory, err := miserCategoryService.GetMiserCategory(uid, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(miserCategory, c)
}

func (miserCategoryApi *MiserCategoryApi) GetMiserCategoryList(c *gin.Context) {
	var pageInfo request.MiserCategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)
	list, total, err := miserCategoryService.GetMiserCategoryInfoList(uid, pageInfo)
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

func (miserCategoryApi *MiserCategoryApi) ListMiserCategoryList(c *gin.Context) {
	uid := utils.GetUserID(c)
	list, err := miserCategoryService.ListMiserCategoryList(uid, nil)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}

	response.OkWithData(list, c)
}
