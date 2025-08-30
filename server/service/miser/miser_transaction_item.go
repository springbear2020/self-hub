package miser

import (
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/miser"
	"github.com/springbear2020/self-hub/server/model/miser/request"
)

type MiserTransactionItemService struct{}

func (miserTransactionItemService *MiserTransactionItemService) CreateMiserTransactionItem(uid uint, miserTransactionItem []miser.MiserTransactionItem) (err error) {
	for i := 0; i < len(miserTransactionItem); i++ {
		miserTransactionItem[i].UserId = uid
	}
	err = global.GVA_DB.Create(miserTransactionItem).Error
	return err
}

func (miserTransactionItemService *MiserTransactionItemService) DeleteMiserTransactionItem(uid uint, id string) (err error) {
	err = global.GVA_DB.Delete(&miser.MiserTransactionItem{}, "user_id = ? AND id = ?", uid, id).Error
	return err
}

func (miserTransactionItemService *MiserTransactionItemService) DeleteMiserTransactionItemByIds(uid uint, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]miser.MiserTransactionItem{}, "user_id = ? AND id in ?", uid, ids).Error
	return err
}

func (miserTransactionItemService *MiserTransactionItemService) UpdateMiserTransactionItem(uid uint, miserTransactionItem miser.MiserTransactionItem) (err error) {
	err = global.GVA_DB.Model(&miser.MiserTransactionItem{}).Where("user_id = ? AND id = ?", uid, miserTransactionItem.Id).Updates(&miserTransactionItem).Error
	return err
}

func (miserTransactionItemService *MiserTransactionItemService) GetMiserTransactionItem(uid uint, id string) (miserTransactionItem miser.MiserTransactionItem, err error) {
	err = global.GVA_DB.Where("user_id = ? AND id = ?", uid, id).First(&miserTransactionItem).Error
	return
}

func (miserTransactionItemService *MiserTransactionItemService) GetMiserTransactionItemInfoList(uid uint, info request.MiserTransactionItemSearch) (list []miser.MiserTransactionItem, total int64, err error) {
	// 查询条件
	db := global.GVA_DB.Model(&miser.MiserTransactionItem{}).Where("user_id = ?", uid)
	if info.TransactionId != nil {
		db = db.Where("transaction_id = ?", *info.TransactionId)
	}
	if info.CategoryId != nil {
		db = db.Where("category_id = ?", *info.CategoryId)
	}
	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
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
	var miserTransactionItems []miser.MiserTransactionItem
	err = db.Order("transaction_id desc, category_id desc, amount desc").Find(&miserTransactionItems).Error
	return miserTransactionItems, total, err
}

func (miserTransactionItemService *MiserTransactionItemService) ListItemDistinctNames(uid uint) (list []string, err error) {
	err = global.GVA_DB.
		Model(&miser.MiserTransactionItem{}).
		Select("DISTINCT name").
		Where("user_id = ?", uid).
		Order("name").
		Find(&list).
		Error
	return
}
