package mine

import (
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/mine"
	"github.com/springbear2020/self-hub/server/model/mine/dto"
	"github.com/springbear2020/self-hub/server/model/mine/request"
)

type MineWebsitesService struct{}

const categoryName = "website_navigation_category"

func (mineWebsitesService *MineWebsitesService) CreateMineWebsites(uid uint, mineWebsites *mine.MineWebsites) (err error) {
	mineWebsites.UserId = uid
	err = global.GVA_DB.Create(mineWebsites).Error
	return err
}

func (mineWebsitesService *MineWebsitesService) DeleteMineWebsites(uid uint, id string) (err error) {
	return mineWebsitesService.DeleteMineWebsitesByIds(uid, []string{id})
}

func (mineWebsitesService *MineWebsitesService) DeleteMineWebsitesByIds(uid uint, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]mine.MineWebsites{}, "user_id = ? AND id in ?", uid, ids).Error
	return err
}

func (mineWebsitesService *MineWebsitesService) UpdateMineWebsites(uid uint, mineWebsites mine.MineWebsites) (err error) {
	err = global.GVA_DB.Model(&mine.MineWebsites{}).Where("user_id = ? AND id = ?", uid, mineWebsites.Id).Updates(&mineWebsites).Error
	return err
}

func (mineWebsitesService *MineWebsitesService) GetMineWebsites(uid uint, id string) (mineWebsites mine.MineWebsites, err error) {
	err = global.GVA_DB.Where("user_id = ? AND id = ?", uid, id).First(&mineWebsites).Error
	return
}

func (mineWebsitesService *MineWebsitesService) GetMineWebsitesInfoList(uid uint, info request.MineWebsitesSearch) (list []mine.MineWebsites, total int64, err error) {
	// 查询条件
	db := global.GVA_DB.Model(&mine.MineWebsites{}).Where("user_id = ? ", uid)
	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	if info.Title != nil && *info.Title != "" {
		db = db.Where("title LIKE ?", "%"+*info.Title+"%")
	}
	if info.Category != nil {
		db = db.Where("category = ?", *info.Category)
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
	var mineWebsites []mine.MineWebsites
	err = db.Order("sort_value desc").Find(&mineWebsites).Error
	return mineWebsites, total, err
}

func (mineWebsitesService *MineWebsitesService) ListMineWebsites(uid uint) (*dto.MineWebsitesDTO, error) {
	websitesDTO := dto.MineWebsitesDTO{}

	// 查询分类字典
	categories, err := dictionaryService.ListAvailable(categoryName)
	if err != nil {
		return nil, err
	}

	websitesDTO.Categories = make([]*dto.MineWebsiteCategory, len(categories))
	for i := range categories {
		websitesDTO.Categories[i] = &dto.MineWebsiteCategory{
			Extend: categories[i].Extend,
			Label:  categories[i].Label,
			Value:  categories[i].Value,
		}
	}

	// 查询网址列表
	var mineWebsites []*mine.MineWebsites
	err = global.GVA_DB.Model(&mine.MineWebsites{}).Where("user_id = ? ", uid).Find(&mineWebsites).Error
	if err != nil {
		return nil, err
	}

	websitesDTO.Sites = make([]*dto.MineWebsite, len(mineWebsites))
	for i := range mineWebsites {
		websitesDTO.Sites[i] = &dto.MineWebsite{
			ID:       mineWebsites[i].Id,
			Name:     mineWebsites[i].Name,
			Title:    mineWebsites[i].Title,
			Href:     mineWebsites[i].Href,
			Icon:     mineWebsites[i].Icon,
			Category: mineWebsites[i].Category,
		}
	}

	return &websitesDTO, nil
}
