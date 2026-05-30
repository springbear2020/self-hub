package example

import (
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/springbear2020/self-hub/server/constants"
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/example/dto"
)

type DashboardService struct{}

func (svc *DashboardService) GetDashboardStatData(uid uint) (dto.DashboardStatDTO, error) {
	var (
		mu sync.Mutex
		g  errgroup.Group
	)

	type statConfig struct {
		Label string
		Table string
		Unit  string
		Icon  string
		Sort  int

		// Count 模式
		IsCount bool

		// Aggregate 模式
		Field string
		Where string
	}

	configs := []statConfig{
		// =========================
		// 聚合统计
		// =========================
		{
			Label: "总收入", Table: "miser_transactions", Unit: "元", Sort: 1, Icon: "Coin",
			Field: "SUM(amount)",
			Where: fmt.Sprintf("user_id = %d AND transaction_type = %d", uid, constants.TransactionTypeIncome),
		},
		{
			Label: "总支出", Table: "miser_transactions", Unit: "元", Sort: 2, Icon: "Money",
			Field: "SUM(amount)",
			Where: fmt.Sprintf("user_id = %d AND transaction_type = %d", uid, constants.TransactionTypeExpense),
		},
		{
			Label: "总结余", Table: "miser_transactions", Unit: "元", Sort: 3, Icon: "Wallet",
			Field: fmt.Sprintf("COALESCE(SUM(IF(transaction_type = %d, amount, 0)), 0) - COALESCE(SUM(IF(transaction_type = %d, amount, 0)), 0)",
				constants.TransactionTypeIncome, constants.TransactionTypeExpense),
			Where: fmt.Sprintf("user_id = %d", uid),
		},
		{
			Label: "总归还", Table: "miser_loan_records", Unit: "元", Sort: 4, Icon: "Coin",
			Field: "SUM(repay_amount)",
			Where: fmt.Sprintf("user_id = %d AND fund_status = %d", uid, constants.FundStatusRepaid),
		},
		{
			Label: "总借出", Table: "miser_loan_records", Unit: "元", Sort: 5, Icon: "Money",
			Field: "SUM(lend_amount)",
			Where: fmt.Sprintf("user_id = %d", uid),
		},
		{
			Label: "待还款", Table: "miser_loan_records", Unit: "元", Sort: 6, Icon: "Wallet",
			Field: "SUM(lend_amount - repay_amount)",
			Where: fmt.Sprintf("user_id = %d AND fund_status != %d", uid, constants.FundStatusRepaid),
		},

		// =========================
		// 总数统计
		// =========================
		{Label: "技术博客", Table: "mine_blogs", Unit: "篇", Sort: 7, IsCount: true, Icon: "Box"},
		{Label: "开源项目", Table: "mine_projects", Unit: "个", Sort: 8, IsCount: true, Icon: "FolderOpened"},
		{Label: "读书札记", Table: "mine_books", Unit: "本", Sort: 9, IsCount: true, Icon: "Notebook"},
		{Label: "电子文藏", Table: "mine_archives", Unit: "项", Sort: 10, IsCount: true, Icon: "Files"},
		{Label: "三五七言", Table: "mine_sentences", Unit: "条", Sort: 11, IsCount: true, Icon: "Comment"},
		{Label: "精品网址", Table: "mine_websites", Unit: "个", Sort: 12, IsCount: true, Icon: "Link"},
		{Label: "今日任务", Table: "daily_tasks", Unit: "个", Sort: 13, IsCount: true, Icon: "Calendar"},
		{Label: "媒体资源", Table: "exa_file_upload_and_downloads", Unit: "项", Sort: 14, IsCount: true, Icon: "Picture"},
	}

	results := make(dto.DashboardStatDTO, 0, len(configs))
	for _, cfg := range configs {
		cfg := cfg
		g.Go(func() error {
			var (
				value string
				err   error
			)

			if cfg.IsCount {
				var total int64
				err = global.GVA_DB.Table(cfg.Table).Where("user_id = ?", uid).Count(&total).Error
				value = fmt.Sprintf("%d", total)
			} else {
				var total float64
				err = global.GVA_DB.Table(cfg.Table).Select(cfg.Field).Where(cfg.Where).Scan(&total).Error
				value = fmt.Sprintf("%.2f", total)
			}

			mu.Lock()
			results = append(results, &dto.DashboardStatItem{
				Label: cfg.Label, Value: value,
				Unit: cfg.Unit, Icon: cfg.Icon,
				Sort: cfg.Sort,
			})
			mu.Unlock()

			return err
		})
	}

	if err := g.Wait(); err != nil {
		return results, err
	}

	results.Ascending()

	return results, nil
}
