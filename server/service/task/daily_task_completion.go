package task

import (
	"fmt"
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/task"
	"github.com/springbear2020/self-hub/server/model/task/request"
	"github.com/springbear2020/self-hub/server/model/task/response"
)

type DailyTaskCompletionService struct{}

var dailyTaskService = new(DailyTaskService)

func (taskCompletionService *DailyTaskCompletionService) CreateDailyTaskCompletion(uid uint, taskCompletion *task.DailyTaskCompletion) (err error) {
	taskIdStr := fmt.Sprintf("%d", *taskCompletion.TaskId)
	dailyTask, err := dailyTaskService.GetDailyTask(uid, taskIdStr)
	if err != nil {
		return err
	}
	if !(*dailyTask.MinValue <= *taskCompletion.CountValue && *taskCompletion.CountValue <= *dailyTask.MaxValue) {
		return fmt.Errorf("任务完成计数值数值范围非法")
	}

	taskCompletion.UserId = uid
	err = global.GVA_DB.Create(taskCompletion).Error
	return err
}

func (taskCompletionService *DailyTaskCompletionService) DeleteDailyTaskCompletion(uid uint, id string) (err error) {
	err = global.GVA_DB.Delete(&task.DailyTaskCompletion{}, "user_id = ? AND id = ?", uid, id).Error
	return err
}

func (taskCompletionService *DailyTaskCompletionService) DeleteDailyTaskCompletionByIds(uid uint, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]task.DailyTaskCompletion{}, "user_id = ? AND id in ?", uid, ids).Error
	return err
}

func (taskCompletionService *DailyTaskCompletionService) UpdateDailyTaskCompletion(uid uint, taskCompletion task.DailyTaskCompletion) (err error) {
	err = global.GVA_DB.Model(&task.DailyTaskCompletion{}).Where("user_id = ? AND id = ?", uid, taskCompletion.Id).Updates(&taskCompletion).Error
	return err
}

func (taskCompletionService *DailyTaskCompletionService) GetDailyTaskCompletion(uid uint, id string) (taskCompletion task.DailyTaskCompletion, err error) {
	err = global.GVA_DB.Where("user_id = ? AND id = ?", uid, id).First(&taskCompletion).Error
	return
}

func (taskCompletionService *DailyTaskCompletionService) GetDailyTaskCompletionInfoList(uid uint, info request.DailyTaskCompletionSearch) (list []task.DailyTaskCompletion, total int64, err error) {
	// 查询条件
	db := global.GVA_DB.Model(&task.DailyTaskCompletion{}).Where("user_id = ? ", uid)
	if info.TaskId != nil && *info.TaskId != 0 {
		db = db.Where("task_id = ?", *info.TaskId)
	}
	if info.FinishDate != nil {
		db = db.Where("finish_date = ?", (*info.FinishDate).Format("2006-01-02"))
	}

	// 查询总数
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

	// 查询列表
	var taskCompletions []task.DailyTaskCompletion
	err = db.Order("finish_date desc").Find(&taskCompletions).Error
	return taskCompletions, total, err
}

func (taskCompletionService *DailyTaskCompletionService) GetDailyTaskStatInfoList(uid uint, req request.DailyTaskStat) (list []response.DailyTaskStat, total int64, err error) {
	// 查询激活任务
	isActive := 1
	taskReq := request.DailyTaskSearch{
		IsActive: &isActive,
		PageInfo: req.PageInfo,
	}
	tasks, total, err := dailyTaskService.GetDailyTaskInfoList(uid, taskReq)
	if err != nil || total == 0 {
		return
	}

	// 遍历任务
	for _, t := range tasks {
		var taskCompletions []task.DailyTaskCompletion
		err = global.GVA_DB.Model(&task.DailyTaskCompletion{}).
			Where("user_id = ? AND task_id = ? AND finish_date BETWEEN ? AND ?",
				uid, t.Id, req.StartDate, req.EndDate).
			Find(&taskCompletions).Error
		if err != nil {
			continue
		}

		completions := make([]response.DailyTaskStatCompletion, len(taskCompletions))
		for i, c := range taskCompletions {
			completions[i].FinishDate = c.FinishDate.Format("2006-01-02")
			completions[i].CountValue = *c.CountValue
		}
		list = append(list, response.DailyTaskStat{
			Task:        t,
			Completions: completions,
		})
	}

	return
}
