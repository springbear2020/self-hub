package miser

import (
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/miser"
	"github.com/springbear2020/self-hub/server/model/miser/request"
)

type MiserTransactionService struct{}

func (miserTransactionService *MiserTransactionService) CreateMiserTransaction(uid uint, miserTransaction []miser.MiserTransaction) (err error) {
	result := miserTransaction[:0] // 重用底层数组
	for _, tx := range miserTransaction {
		if tx.Amount != nil && *tx.Amount != 0 {
			tx.UserId = uid
			result = append(result, tx)
		}
	}
	err = global.GVA_DB.Create(result).Error
	return err
}

func (miserTransactionService *MiserTransactionService) DeleteMiserTransaction(uid uint, id string) (err error) {
	err = global.GVA_DB.Delete(&miser.MiserTransaction{}, "user_id = ? AND id = ?", uid, id).Error
	return err
}

func (miserTransactionService *MiserTransactionService) DeleteMiserTransactionByIds(uid uint, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]miser.MiserTransaction{}, "user_id = ? AND id in ?", uid, ids).Error
	return err
}

func (miserTransactionService *MiserTransactionService) UpdateMiserTransaction(uid uint, miserTransaction miser.MiserTransaction) (err error) {
	err = global.GVA_DB.Model(&miser.MiserTransaction{}).Where("user_id = ? AND id = ?", uid, miserTransaction.Id).Updates(&miserTransaction).Error
	return err
}

func (miserTransactionService *MiserTransactionService) GetMiserTransaction(uid uint, id string) (miserTransaction miser.MiserTransaction, err error) {
	err = global.GVA_DB.Where("user_id = ? AND id = ?", uid, id).First(&miserTransaction).Error
	return
}

func (miserTransactionService *MiserTransactionService) GetMiserTransactionInfoList(uid uint, info request.MiserTransactionSearch) (list []miser.MiserTransaction, total int64, err error) {
	// 查询条件
	db := global.GVA_DB.Model(&miser.MiserTransaction{}).Where("user_id = ?", uid)
	if info.CategoryId != nil {
		db = db.Where("category_id = ?", *info.CategoryId)
	}
	if info.TransactionType != nil {
		db = db.Where("transaction_type = ?", *info.TransactionType)
	}
	if info.Date != nil {
		db = db.Where("date = ?", *info.Date)
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
	var miserTransactions []miser.MiserTransaction
	err = db.Order("date desc, transaction_type desc, amount desc").Find(&miserTransactions).Error
	return miserTransactions, total, err
}
