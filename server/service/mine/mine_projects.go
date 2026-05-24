package mine

import (
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/mine"
	"github.com/springbear2020/self-hub/server/model/mine/dto"
	"github.com/springbear2020/self-hub/server/model/mine/request"
)

type MineProjectsService struct{}

func (mineProjectsService *MineProjectsService) CreateMineProjects(uid uint, mineProjects *mine.MineProjects) (err error) {
	mineProjects.UserId = uid
	err = global.GVA_DB.Create(mineProjects).Error
	return err
}

func (mineProjectsService *MineProjectsService) DeleteMineProjects(uid uint, id string) (err error) {
	return mineProjectsService.DeleteMineProjectsByIds(uid, []string{id})
}

func (mineProjectsService *MineProjectsService) DeleteMineProjectsByIds(uid uint, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]mine.MineProjects{}, "user_id = ? AND id in ?", uid, ids).Error
	return err
}

func (mineProjectsService *MineProjectsService) UpdateMineProjects(uid uint, mineProjects mine.MineProjects) (err error) {
	err = global.GVA_DB.Model(&mine.MineProjects{}).Where("user_id = ? AND id = ?", uid, mineProjects.Id).Updates(&mineProjects).Error
	return err
}

func (mineProjectsService *MineProjectsService) GetMineProjects(uid uint, id string) (mineProjects mine.MineProjects, err error) {
	err = global.GVA_DB.Where("user_id = ? AND id = ?", uid, id).First(&mineProjects).Error
	return
}

func (mineProjectsService *MineProjectsService) GetMineProjectsInfoList(uid uint, info request.MineProjectsSearch) (list []mine.MineProjects, total int64, err error) {
	// 查询条件
	db := global.GVA_DB.Model(&mine.MineProjects{}).Where("user_id = ? ", uid)
	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	if info.Title != nil && *info.Title != "" {
		db = db.Where("title LIKE ?", "%"+*info.Title+"%")
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
	var mineProjects []mine.MineProjects
	err = db.Order("sort_value desc, start_date desc").Find(&mineProjects).Error
	return mineProjects, total, err
}

func (mineProjectsService *MineProjectsService) SearchMineProjects(uid uint, keyword string) (list []*dto.MineResourcesDTO, err error) {
	db := global.GVA_DB.
		Table("mine_projects").
		Model(&dto.MineResourcesDTO{}).
		Where("user_id = ? AND (title LIKE ? OR name LIKE ?)", uid, "%"+keyword+"%", "%"+keyword+"%")
	err = db.Select("id, title, name as description, link").Order("sort_value desc, start_date desc").Find(&list).Limit(5).Error
	if err != nil {
		return
	}

	for _, item := range list {
		item.Type = "project"
	}
	return
}
