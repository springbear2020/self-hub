package miser

import (
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/miser"
	"github.com/springbear2020/self-hub/server/model/miser/request"
)

type MiserRankingRecordService struct{}

func (miserRankingRecordService *MiserRankingRecordService) CreateMiserRankingRecord(uid uint, miserRankingRecord *miser.MiserRankingRecord) (err error) {
	miserRankingRecord.UserId = uid
	err = global.GVA_DB.Create(miserRankingRecord).Error
	return err
}

func (miserRankingRecordService *MiserRankingRecordService) DeleteMiserRankingRecord(uid uint, id string) (err error) {
	err = global.GVA_DB.Delete(&miser.MiserRankingRecord{}, "user_id = ? AND id = ?", uid, id).Error
	return err
}

func (miserRankingRecordService *MiserRankingRecordService) DeleteMiserRankingRecordByIds(uid uint, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]miser.MiserRankingRecord{}, "user_id = ? AND id in ?", uid, ids).Error
	return err
}

func (miserRankingRecordService *MiserRankingRecordService) UpdateMiserRankingRecord(uid uint, miserRankingRecord miser.MiserRankingRecord) (err error) {
	err = global.GVA_DB.Model(&miser.MiserRankingRecord{}).Where("user_id = ? AND id = ?", uid, miserRankingRecord.Id).Updates(&miserRankingRecord).Error
	return err
}

func (miserRankingRecordService *MiserRankingRecordService) GetMiserRankingRecord(uid uint, id string) (miserRankingRecord miser.MiserRankingRecord, err error) {
	err = global.GVA_DB.Where("user_id = ? AND id = ?", uid, id).First(&miserRankingRecord).Error
	return
}

func (miserRankingRecordService *MiserRankingRecordService) GetMiserRankingRecordInfoList(uid uint, info request.MiserRankingRecordSearch) (list []miser.MiserRankingRecord, total int64, err error) {
	// 查询条件
	db := global.GVA_DB.Model(&miser.MiserRankingRecord{}).Where("user_id = ?", uid)
	if info.TransactionType != nil {
		db = db.Where("transaction_type = ?", *info.TransactionType)
	}
	if info.CategoryId != nil {
		db = db.Where("category_id = ?", *info.CategoryId)
	}
	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
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
	var miserRankingRecords []miser.MiserRankingRecord
	err = db.Order("date desc, amount desc").Find(&miserRankingRecords).Error
	return miserRankingRecords, total, err
}
