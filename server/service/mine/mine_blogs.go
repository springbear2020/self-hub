package mine

import (
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/mine"
	"github.com/springbear2020/self-hub/server/model/mine/dto"
	"github.com/springbear2020/self-hub/server/model/mine/request"
)

type MineBlogsService struct{}

func (mineBlogsService *MineBlogsService) CreateMineBlogs(uid uint, mineBlogs *mine.MineBlogs) (err error) {
	mineBlogs.UserId = uid
	err = global.GVA_DB.Create(mineBlogs).Error
	return err
}

func (mineBlogsService *MineBlogsService) DeleteMineBlogs(uid uint, id string) (err error) {
	return mineBlogsService.DeleteMineBlogsByIds(uid, []string{id})
}

func (mineBlogsService *MineBlogsService) DeleteMineBlogsByIds(uid uint, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]mine.MineBlogs{}, "user_id = ? AND id in ?", uid, ids).Error
	return err
}

func (mineBlogsService *MineBlogsService) UpdateMineBlogs(uid uint, mineBlogs mine.MineBlogs) (err error) {
	err = global.GVA_DB.Model(&mine.MineBlogs{}).Where("user_id = ? AND id = ?", uid, mineBlogs.Id).Updates(&mineBlogs).Error
	return err
}

func (mineBlogsService *MineBlogsService) GetMineBlogs(uid uint, id string) (mineBlogs mine.MineBlogs, err error) {
	err = global.GVA_DB.Where("user_id = ? AND id = ?", uid, id).First(&mineBlogs).Error
	return
}

func (mineBlogsService *MineBlogsService) GetMineBlogsInfoList(uid uint, info request.MineBlogsSearch) (list []mine.MineBlogs, total int64, err error) {
	// 查询条件
	db := global.GVA_DB.Model(&mine.MineBlogs{}).Where("user_id = ? ", uid)
	if info.Title != nil && *info.Title != "" {
		db = db.Where("title LIKE ?", "%"+*info.Title+"%")
	}
	if info.CategoryId != nil {
		db = db.Where("category_id = ?", *info.CategoryId)
	}
	if info.TagId != nil {
		db = db.Where("tag_id = ?", *info.TagId)
	}

	// 记录总数
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// 分页条件
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	// 查询记录
	var mineBlogs []mine.MineBlogs
	err = db.Order("sort_value desc, post_time desc").Find(&mineBlogs).Error
	return mineBlogs, total, err
}

func (mineBlogsService *MineBlogsService) SearchMineBlogs(uid uint, keyword string) (list []*dto.MineResourcesDTO, err error) {
	db := global.GVA_DB.
		Table("mine_blogs").
		Model(&dto.MineResourcesDTO{}).
		Where("user_id = ? AND title LIKE ?", uid, "%"+keyword+"%")
	err = db.Select("id, title, DATE_FORMAT(post_time, '%Y-%m-%d %H:%i') as description, link").Order("sort_value desc, post_time desc").Find(&list).Limit(5).Error
	if err != nil {
		return
	}

	for _, item := range list {
		item.Type = "blog"
	}
	return
}
