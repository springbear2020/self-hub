package miser

import (
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/miser"
	"github.com/springbear2020/self-hub/server/model/miser/request"
	"github.com/springbear2020/self-hub/server/utils"
)

type MiserLoanRecordService struct{}

func (miserLoanRecordService *MiserLoanRecordService) CreateMiserLoanRecord(uid uint, miserLoanRecord *miser.MiserLoanRecord) (err error) {
	if err = miserLoanRecord.CalculateLoanRecord(); err != nil {
		return err
	}
	miserLoanRecord.UserId = uid
	err = global.GVA_DB.Create(miserLoanRecord).Error
	return err
}

func (miserLoanRecordService *MiserLoanRecordService) DeleteMiserLoanRecord(uid uint, id string) (err error) {
	err = global.GVA_DB.Delete(&miser.MiserLoanRecord{}, "user_id = ? AND id = ?", uid, id).Error
	return err
}

func (miserLoanRecordService *MiserLoanRecordService) DeleteMiserLoanRecordByIds(uid uint, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]miser.MiserLoanRecord{}, "user_id = ? AND id in ?", uid, ids).Error
	return err
}

func (miserLoanRecordService *MiserLoanRecordService) UpdateMiserLoanRecord(uid uint, miserLoanRecord miser.MiserLoanRecord) (err error) {
	if err = miserLoanRecord.CalculateLoanRecord(); err != nil {
		return err
	}
	err = global.GVA_DB.Model(&miser.MiserLoanRecord{}).Where("user_id = ? AND id = ?", uid, miserLoanRecord.Id).Updates(&miserLoanRecord).Error
	return err
}

func (miserLoanRecordService *MiserLoanRecordService) GetMiserLoanRecord(uid uint, id string) (miserLoanRecord miser.MiserLoanRecord, err error) {
	err = global.GVA_DB.Where("user_id = ? AND id = ?", uid, id).First(&miserLoanRecord).Error
	return
}

func (miserLoanRecordService *MiserLoanRecordService) GetMiserLoanRecordInfoList(uid uint, info request.MiserLoanRecordSearch) (list []miser.MiserLoanRecord, total int64, err error) {
	// 查询条件
	db := global.GVA_DB.Model(&miser.MiserLoanRecord{}).Where("user_id = ?", uid)
	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	if info.LendDate != nil {
		db = db.Where("lend_date = ?", *info.LendDate)
	}
	if info.RepayDate != nil {
		db = db.Where("repay_date = ?", *info.RepayDate)
	}
	if info.FundStatus != nil {
		db = db.Where("fund_status = ?", *info.FundStatus)
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
	var miserLoanRecords []miser.MiserLoanRecord
	err = db.Order("lend_date desc, fund_status").Find(&miserLoanRecords).Error
	return miserLoanRecords, total, err
}

func (miserLoanRecordService *MiserLoanRecordService) ListMiserLoanNameList(uid uint) (list []string, err error) {
	db := global.GVA_DB.
		Model(&miser.MiserLoanRecord{}).
		Select("DISTINCT name").
		Where("user_id = ?", uid)

	err = db.Find(&list).Error
	if err != nil {
		return
	}

	// 中文名排序
	utils.SortChineseNames(list)

	return
}

func (miserLoanRecordService *MiserLoanRecordService) GetMiserLoanStatData(uid uint) (interface{}, error) {
	var result struct {
		LendAmount     float64 `json:"lend_amount"`     // 总借出
		RepaidAmount   float64 `json:"repaid_amount"`   // 总归还
		RepayingAmount float64 `json:"repaying_amount"` // 待还款
	}

	err := global.GVA_DB.Raw(`
		SELECT SUM(lend_amount)                       AS lend_amount,
		       SUM(repay_amount)                      AS repaid_amount,
		       (SUM(lend_amount) - SUM(repay_amount)) AS repaying_amount
		FROM miser_loan_records
		WHERE user_id = ?;
	`, uid).Scan(&result).Error

	return result, err
}
