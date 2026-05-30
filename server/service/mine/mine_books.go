package mine

import (
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/mine"
	"github.com/springbear2020/self-hub/server/model/mine/dto"
	"github.com/springbear2020/self-hub/server/model/mine/request"
)

type MineBooksService struct{}

func (mineBooksService *MineBooksService) CreateMineBooks(uid uint, mineBooks *mine.MineBooks) (err error) {
	mineBooks.UserId = uid
	err = global.GVA_DB.Create(mineBooks).Error
	return err
}

func (mineBooksService *MineBooksService) DeleteMineBooks(uid uint, id string) (err error) {
	return mineBooksService.DeleteMineBooksByIds(uid, []string{id})
}

func (mineBooksService *MineBooksService) DeleteMineBooksByIds(uid uint, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]mine.MineBooks{}, "user_id = ? AND id in ?", uid, ids).Error
	return err
}

func (mineBooksService *MineBooksService) UpdateMineBooks(uid uint, mineBooks mine.MineBooks) (err error) {
	err = global.GVA_DB.Model(&mine.MineBooks{}).Where("user_id = ? AND id = ?", uid, mineBooks.Id).Updates(&mineBooks).Error
	return err
}

func (mineBooksService *MineBooksService) GetMineBooks(uid uint, id string) (mineBooks mine.MineBooks, err error) {
	err = global.GVA_DB.Where("user_id = ? AND id = ?", uid, id).First(&mineBooks).Error
	return
}

func (mineBooksService *MineBooksService) GetMineBooksInfoList(uid uint, info request.MineBooksSearch) (list []mine.MineBooks, total int64, err error) {
	// 查询条件
	db := global.GVA_DB.Model(&mine.MineBooks{}).Where("user_id = ? ", uid)
	if info.Title != nil && *info.Title != "" {
		db = db.Where("title LIKE ?", "%"+*info.Title+"%")
	}
	if info.Author != nil && *info.Author != "" {
		db = db.Where("author LIKE ?", "%"+*info.Author+"%")
	}
	if info.IconicQuote != nil && *info.IconicQuote != "" {
		db = db.Where("iconic_quote LIKE ?", "%"+*info.IconicQuote+"%")
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
	var mineBooks []mine.MineBooks
	err = db.Order("sort_value desc, first_completion desc").Find(&mineBooks).Error
	return mineBooks, total, err
}

func (mineBooksService *MineBooksService) SearchMineBooks(uid uint, keyword string) (list []*dto.MineResourcesDTO, err error) {
	db := global.GVA_DB.
		Table("mine_books").
		Model(&dto.MineResourcesDTO{}).
		Where("user_id = ? AND title LIKE ?", uid, "%"+keyword+"%")
	err = db.Select("id, title, author as description, link").Order("sort_value desc, first_completion desc").Limit(searchPageSize).Find(&list).Error
	if err != nil {
		return
	}

	for _, item := range list {
		item.Type = searchTypeBook
	}
	return
}
