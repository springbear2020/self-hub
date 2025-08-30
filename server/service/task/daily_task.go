package task

import (
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/task"
	"github.com/springbear2020/self-hub/server/model/task/request"
)

type DailyTaskService struct{}

func (dailyTaskService *DailyTaskService) CreateDailyTask(uid uint, dailyTask *task.DailyTask) (err error) {
	dailyTask.UserId = uid
	err = global.GVA_DB.Create(dailyTask).Error
	return err
}

func (dailyTaskService *DailyTaskService) DeleteDailyTask(uid uint, id string) (err error) {
	err = global.GVA_DB.Delete(&task.DailyTask{}, "user_id = ? AND id = ?", uid, id).Error
	return err
}

func (dailyTaskService *DailyTaskService) DeleteDailyTaskByIds(uid uint, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]task.DailyTask{}, "user_id = ? AND id in ?", uid, ids).Error
	return err
}

func (dailyTaskService *DailyTaskService) UpdateDailyTask(uid uint, dailyTask task.DailyTask) (err error) {
	err = global.GVA_DB.Model(&task.DailyTask{}).Where("user_id = ? AND id = ?", uid, dailyTask.Id).Updates(&dailyTask).Error
	return err
}

func (dailyTaskService *DailyTaskService) GetDailyTask(uid uint, id string) (dailyTask task.DailyTask, err error) {
	err = global.GVA_DB.Where("user_id = ? AND id = ?", uid, id).First(&dailyTask).Error
	return
}

func (dailyTaskService *DailyTaskService) GetDailyTaskInfoList(uid uint, info request.DailyTaskSearch) (list []task.DailyTask, total int64, err error) {
	// 查询条件
	db := global.GVA_DB.Model(&task.DailyTask{}).Where("user_id = ?", uid)
	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	if info.IsActive != nil {
		db = db.Where("is_active = ?", *info.IsActive)
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
	var dailyTasks []task.DailyTask
	err = db.Order("sort desc").Find(&dailyTasks).Error
	return dailyTasks, total, err
}

func (dailyTaskService *DailyTaskService) ListActiveTaskList(uid uint) (list []task.DailyTask, err error) {
	err = global.GVA_DB.Where("user_id = ? AND is_active = 1", uid).Order("sort desc").Find(&list).Error
	return
}
