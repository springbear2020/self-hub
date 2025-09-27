package miser

import (
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/miser"
	"github.com/springbear2020/self-hub/server/model/miser/request"
	"gorm.io/gorm"
)

type MiserCategoryService struct{}

func (miserCategoryService *MiserCategoryService) CreateMiserCategory(uid uint, miserCategory *miser.MiserCategory) (err error) {
	miserCategory.UserId = uid
	err = global.GVA_DB.Create(miserCategory).Error
	return err
}

func (miserCategoryService *MiserCategoryService) DeleteMiserCategory(uid uint, id string) (err error) {
	return miserCategoryService.DeleteMiserCategoryByIds(uid, []string{id})
}

func (miserCategoryService *MiserCategoryService) DeleteMiserCategoryByIds(uid uint, ids []string) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 删除交易分类时删除交易流水和流水明细
		iErr := tx.Delete(&miser.MiserCategory{}, "user_id = ? AND id IN ?", uid, ids).Error
		if iErr != nil {
			return iErr
		}
		iErr = tx.Delete(&miser.MiserTransaction{}, "user_id = ? AND category_id IN ?", uid, ids).Error
		if iErr != nil {
			return iErr
		}
		return tx.Delete(&miser.MiserTransactionItem{}, "user_id = ? AND category_id IN ?", uid, ids).Error
	})
	return err
}

func (miserCategoryService *MiserCategoryService) UpdateMiserCategory(uid uint, miserCategory miser.MiserCategory) (err error) {
	err = global.GVA_DB.Model(&miser.MiserCategory{}).Where("user_id = ? AND id = ?", uid, miserCategory.Id).Updates(&miserCategory).Error
	return err
}

func (miserCategoryService *MiserCategoryService) GetMiserCategory(uid uint, id string) (miserCategory miser.MiserCategory, err error) {
	err = global.GVA_DB.Where("user_id = ? AND id = ?", uid, id).First(&miserCategory).Error
	return
}

func (miserCategoryService *MiserCategoryService) GetMiserCategoryInfoList(uid uint, info request.MiserCategorySearch) (list []miser.MiserCategory, total int64, err error) {
	// 查询条件
	db := global.GVA_DB.Model(&miser.MiserCategory{}).Where("user_id = ? ", uid)
	if info.TransactionType != nil {
		db = db.Where("transaction_type = ?", *info.TransactionType)
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
	var miserCategories []miser.MiserCategory
	err = db.Order("sort desc").Find(&miserCategories).Error
	return miserCategories, total, err
}

func (miserCategoryService *MiserCategoryService) ListMiserCategoryList(uid uint, transactionType *int) (list []miser.MiserCategory, err error) {
	db := global.GVA_DB.Model(&miser.MiserCategory{}).Where("user_id = ?", uid)
	if transactionType != nil {
		db = db.Where("transaction_type = ?", *transactionType)
	}
	err = db.Order("sort desc").Find(&list).Error
	return
}
