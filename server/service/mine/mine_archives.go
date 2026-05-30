package mine

import (
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/mine"
	"github.com/springbear2020/self-hub/server/model/mine/dto"
	"github.com/springbear2020/self-hub/server/model/mine/request"
)

type MineArchivesService struct{}

func (mineArchivesService *MineArchivesService) CreateMineArchives(uid uint, mineArchives *mine.MineArchives) (err error) {
	mineArchives.UserId = uid
	err = global.GVA_DB.Create(mineArchives).Error
	return err
}

func (mineArchivesService *MineArchivesService) DeleteMineArchives(uid uint, id string) (err error) {
	return mineArchivesService.DeleteMineArchivesByIds(uid, []string{id})
}

func (mineArchivesService *MineArchivesService) DeleteMineArchivesByIds(uid uint, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]mine.MineArchives{}, "user_id = ? AND id in ?", uid, ids).Error
	return err
}

func (mineArchivesService *MineArchivesService) UpdateMineArchives(uid uint, mineArchives mine.MineArchives) (err error) {
	err = global.GVA_DB.Model(&mine.MineArchives{}).Where("user_id = ? AND id = ?", uid, mineArchives.Id).Updates(&mineArchives).Error
	return err
}

func (mineArchivesService *MineArchivesService) GetMineArchives(uid uint, id string) (mineArchives mine.MineArchives, err error) {
	err = global.GVA_DB.Where("user_id = ? AND id = ?", uid, id).First(&mineArchives).Error
	return
}

func (mineArchivesService *MineArchivesService) GetMineArchivesInfoList(uid uint, info request.MineArchivesSearch) (list []mine.MineArchives, total int64, err error) {
	// 查询条件
	db := global.GVA_DB.Model(&mine.MineArchives{}).Where("user_id = ? ", uid)
	if info.Title != nil && *info.Title != "" {
		db = db.Where("title LIKE ?", "%"+*info.Title+"%")
	}
	if info.Category != nil && *info.Category != "" {
		db = db.Where("category LIKE ?", "%"+*info.Category+"%")
	}
	if info.Description != nil && *info.Description != "" {
		db = db.Where("description LIKE ?", "%"+*info.Description+"%")
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
	var mineArchives []mine.MineArchives
	err = db.Order("sort_value desc").Find(&mineArchives).Error
	return mineArchives, total, err
}

func (mineArchivesService *MineArchivesService) SearchMineArchives(uid uint, keyword string) (list []*dto.MineResourcesDTO, err error) {
	db := global.GVA_DB.
		Table("mine_archives").
		Model(&dto.MineResourcesDTO{}).
		Where("user_id = ? AND (title LIKE ? OR description LIKE ?)", uid, "%"+keyword+"%", "%"+keyword+"%")
	err = db.Select("id, title, description, link").Order("sort_value desc").Limit(searchPageSize).Find(&list).Error
	if err != nil {
		return
	}

	for _, item := range list {
		item.Type = searchTypeArchive
	}
	return
}
