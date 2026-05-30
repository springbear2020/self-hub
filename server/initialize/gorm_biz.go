package initialize

import (
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/mine"
	"github.com/springbear2020/self-hub/server/model/miser"
	"github.com/springbear2020/self-hub/server/model/task"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(
		task.DailyTask{}, task.DailyTaskCompletion{},
		miser.MiserCategory{}, miser.MiserTransaction{}, miser.MiserTransactionItem{}, miser.MiserLoanRecord{}, miser.MiserRankingRecord{},
		mine.MineBooks{}, mine.MineBlogs{}, mine.MineProjects{}, mine.MineWebsites{}, mine.MineArchives{}, mine.MineSentences{},
	)
	if err != nil {
		return err
	}
	return nil
}
