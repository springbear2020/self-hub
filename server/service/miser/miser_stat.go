package miser

import (
	"fmt"
	"github.com/springbear2020/self-hub/server/constants"
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/miser/dto"
	"github.com/springbear2020/self-hub/server/model/miser/request"
	"golang.org/x/sync/errgroup"
	"math"
	"strconv"
	"strings"
)

type MiserStatService struct{}

var miserCategoryService = MiserCategoryService{}

func (s *MiserStatService) GetRankingStat(uid uint, req *request.MiserCategoryStat) (map[int][]dto.MiserStatRankingAmount, error) {
	var g errgroup.Group

	var (
		incomes  []dto.MiserStatRankingAmount
		expenses []dto.MiserStatRankingAmount
	)

	g.Go(func() error {
		var err error
		incomes, err = s.getRankingAmount(uid, constants.TransactionTypeIncome, req)
		return err
	})
	g.Go(func() error {
		var err error
		expenses, err = s.getRankingAmount(uid, constants.TransactionTypeExpense, req)
		return err
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return map[int][]dto.MiserStatRankingAmount{
		constants.TransactionTypeIncome:  incomes,
		constants.TransactionTypeExpense: expenses,
	}, nil
}

func (s *MiserStatService) GetCategoryItemStat(uid uint, req *request.MiserCategoryItemStat) ([]dto.MiserStatMonthAmount, error) {
	var results []dto.MiserStatMonthAmount
	err := global.GVA_DB.Raw(`
		SELECT DATE_FORMAT(date, '%Y-%m') AS month,
			   amount                     AS amount
		FROM miser_transaction_items
		WHERE user_id = ?
		  AND category_id = ?
		  AND name = ?
		  AND DATE_FORMAT(date, '%Y-%m') BETWEEN ? AND ?
		ORDER BY month;
	`, uid, req.CategoryId, req.Name, req.StartMonth, req.EndMonth,
	).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *MiserStatService) GetMonthTransactionStat(uid uint, req *request.MiserTransactionStat) ([]dto.MiserStatCategoryAmount, error) {
	miserStatReq := &request.MiserStat{
		StartMonth: req.StartMonth,
		EndMonth:   req.EndMonth,
	}
	return s.getCategoryAmount(uid, req.TransactionType, miserStatReq)
}

func (s *MiserStatService) GetCategoryStat(uid uint, req *request.MiserCategoryStat) ([]dto.MiserStatMonthAmount, error) {
	var results []dto.MiserStatMonthAmount
	err := global.GVA_DB.Raw(`
		SELECT DATE_FORMAT(date, '%Y-%m') AS month,
			   amount                     AS amount
		FROM miser_transactions
		WHERE user_id = ?
		  AND category_id = ?
		  AND DATE_FORMAT(date, '%Y-%m') BETWEEN ? AND ?
		ORDER BY date;
	`, uid, req.CategoryId, req.StartMonth, req.EndMonth,
	).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *MiserStatService) GetItemStat(uid uint, req *request.MiserStat) (map[int][]dto.MiserStatTransactionAmount, error) {
	var results []dto.MiserStatTransactionAmount
	err := global.GVA_DB.Raw(`
	SELECT category_id, name, SUM(amount) AS amount
	FROM miser_transaction_items
	WHERE user_id = ?
	  AND transaction_id IN (SELECT DISTINCT id
							 FROM miser_transactions
							 WHERE user_id = ?
							   AND DATE_FORMAT(date, '%Y-%m') BETWEEN ? AND ?)
	GROUP BY category_id, name
	ORDER BY category_id, amount DESC;
	`, uid, uid, req.StartMonth, req.EndMonth,
	).Find(&results).Error
	if err != nil {
		return nil, err
	}

	// 将列表中的数据根据 category_id 分组为 map 后返回
	grouped := make(map[int][]dto.MiserStatTransactionAmount)
	for _, r := range results {
		grouped[r.CategoryId] = append(grouped[r.CategoryId], r)
	}

	return grouped, nil
}

func (s *MiserStatService) GetYearStat(uid uint, req *request.MiserStat) ([]dto.MiserStatYear, error) {
	var results []dto.MiserStatYear
	err := global.GVA_DB.Raw(`
	SELECT DATE_FORMAT(date, '%Y')                                                               AS year,
		   SUM(IF(transaction_type = ?, amount, 0))                                              AS income,
		   SUM(IF(transaction_type = ?, amount, 0))                                              AS expense,
		   (SUM(IF(transaction_type = ?, amount, 0)) - SUM(IF(transaction_type = ?, amount, 0))) AS balance
	FROM miser_transactions
	WHERE user_id = ?
	  AND DATE_FORMAT(date, '%Y') BETWEEN ? AND ?
	GROUP BY year
	ORDER BY year;
	`, constants.TransactionTypeIncome, constants.TransactionTypeExpense, constants.TransactionTypeIncome, constants.TransactionTypeExpense, uid, req.StartMonth[:4], req.EndMonth[:4],
	).Find(&results).Error
	return results, err
}

func (s *MiserStatService) GetMonthStat(uid uint, req *request.MiserStat) ([]dto.TableData, error) {
	var g errgroup.Group
	income := dto.TableData{HeaderId: constants.TransactionTypeIncome, Header: "收入明细"}
	expense := dto.TableData{HeaderId: constants.TransactionTypeExpense, Header: "支出明细"}

	g.Go(func() error {
		var err error
		income.Rows, income.Cols, err = s.getCategoryTableData(uid, req, constants.TransactionTypeIncome)
		return err
	})

	g.Go(func() error {
		var err error
		expense.Rows, expense.Cols, err = s.getCategoryTableData(uid, req, constants.TransactionTypeExpense)
		return err
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return []dto.TableData{income, expense}, nil
}

func (s *MiserStatService) GetLineStat(uid uint, req *request.MiserStat) ([]dto.MiserStatMonth, error) {
	var results []dto.MiserStatMonth
	err := global.GVA_DB.Raw(`
	SELECT DATE_FORMAT(date, '%Y-%m')                                                            AS month,
		   SUM(IF(transaction_type = ?, amount, 0))                                              AS income,
		   SUM(IF(transaction_type = ?, amount, 0))                                              AS expense,
		   (SUM(IF(transaction_type = ?, amount, 0)) - SUM(IF(transaction_type = ?, amount, 0))) AS balance
	FROM miser_transactions
	WHERE user_id = ?
	  AND DATE_FORMAT(date, '%Y-%m') BETWEEN ? AND ?
	GROUP BY month
	ORDER BY month;
	`, constants.TransactionTypeIncome, constants.TransactionTypeExpense, constants.TransactionTypeIncome, constants.TransactionTypeExpense, uid, req.StartMonth, req.EndMonth,
	).Find(&results).Error
	return results, err
}

func (s *MiserStatService) GetPieStat(uid uint, req *request.MiserStat) (map[string][]dto.MiserStatCategoryAmount, error) {
	var g errgroup.Group
	var incomes, expenses []dto.MiserStatCategoryAmount

	g.Go(func() error {
		var err error
		incomes, err = s.getCategoryAmount(uid, constants.TransactionTypeIncome, req)
		return err
	})

	g.Go(func() error {
		var err error
		expenses, err = s.getCategoryAmount(uid, constants.TransactionTypeExpense, req)
		return err
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return map[string][]dto.MiserStatCategoryAmount{
		"incomes":  incomes,
		"expenses": expenses,
	}, nil
}

func (s *MiserStatService) GetCardStat(uid uint, req *request.MiserStat) (map[string]float64, error) {
	var g errgroup.Group
	var totalIncome, totalExpense float64

	g.Go(func() error {
		var err error
		totalIncome, err = s.getTotalAmount(uid, constants.TransactionTypeIncome, req)
		return err
	})

	g.Go(func() error {
		var err error
		totalExpense, err = s.getTotalAmount(uid, constants.TransactionTypeExpense, req)
		return err
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	totalBalance := math.Round((totalIncome-totalExpense)*100) / 100
	return map[string]float64{
		"income":  totalIncome,
		"expense": totalExpense,
		"balance": totalBalance,
	}, nil
}

func (s *MiserStatService) getCategoryTableData(uid uint, req *request.MiserStat, transactionType int) ([]dto.TableRow, []dto.TableColumn, error) {
	// 查询分类列表
	categories, err := miserCategoryService.ListMiserCategoryList(uid, &transactionType)
	if err != nil {
		return nil, nil, err
	}

	// 构建动态列名
	selectFields := []string{"DATE_FORMAT(date, '%Y-%m') AS month"}
	for _, c := range categories {
		field := fmt.Sprintf("SUM(IF(category_id = %d, amount, 0)) AS `%s`", *c.Id, *c.Name)
		selectFields = append(selectFields, field)
	}
	sql := fmt.Sprintf(`
		SELECT %s
		FROM miser_transactions
		WHERE user_id = ? 
		  AND transaction_type = ?
		  AND DATE_FORMAT(date, '%%Y-%%m') BETWEEN ? AND ?
		GROUP BY month
		ORDER BY month DESC
	`, strings.Join(selectFields, ", "))
	var results []map[string]interface{}
	err = global.GVA_DB.Raw(sql, uid, transactionType, req.StartMonth, req.EndMonth).Scan(&results).Error
	if err != nil {
		return nil, nil, err
	}

	// 行数据
	rows := make([]dto.TableRow, len(results))
	for i, row := range results {
		// 处理月份字段
		if month, ok := row["month"]; ok {
			row["月份"] = month
			delete(row, "month")
		}

		// 处理金额字段
		var total float64
		for key, value := range row {
			if f, iErr := strconv.ParseFloat(value.(string), 64); iErr == nil {
				row[key] = f
				total += f
			}
		}
		row["金额"] = math.Round(total*100) / 100

		rows[i] = row
	}

	// 列表头
	cols := []dto.TableColumn{
		{Label: "月份", Key: "月份", Fixed: true, Sortable: true},
		{Label: "金额", Key: "金额", Fixed: true, Sortable: true},
	}
	for _, c := range categories {
		cols = append(cols, dto.TableColumn{
			Label: *c.Name,
			Key:   *c.Name,
		})
	}

	return rows, cols, nil
}

func (s *MiserStatService) getCategoryAmount(uid uint, transactionType int, req *request.MiserStat) ([]dto.MiserStatCategoryAmount, error) {
	var results []dto.MiserStatCategoryAmount
	err := global.GVA_DB.Raw(`
	SELECT category_id AS category, SUM(amount) AS amount
	FROM miser_transactions
	WHERE user_id = ?
	  AND transaction_type = ?
	  AND DATE_FORMAT(date, '%Y-%m') BETWEEN ? AND ?
	GROUP BY category_id
	ORDER BY amount DESC;
	`, uid, transactionType, req.StartMonth, req.EndMonth,
	).Find(&results).Error
	return results, err
}

func (s *MiserStatService) getTotalAmount(uid uint, transactionType int, req *request.MiserStat) (float64, error) {
	var amount float64
	err := global.GVA_DB.Raw(`
	SELECT IFNULL(SUM(amount), 0) AS amount
	FROM miser_transactions
	WHERE user_id = ?
	  AND transaction_type = ?
	  AND DATE_FORMAT(date, '%Y-%m') BETWEEN ? AND ?;
	`, uid, transactionType, req.StartMonth, req.EndMonth,
	).Scan(&amount).Error
	return amount, err
}

func (s *MiserStatService) getRankingAmount(uid uint, transactionType int, req *request.MiserCategoryStat) ([]dto.MiserStatRankingAmount, error) {
	var results []dto.MiserStatRankingAmount

	// 基础 SQL
	query := `
	SELECT category_id, name, amount, date
	FROM miser_ranking_records
	WHERE user_id = ?
	  AND transaction_type = ?
	  AND DATE_FORMAT(date, '%Y-%m') BETWEEN ? AND ?
	`

	// 动态参数列表
	args := []interface{}{uid, transactionType, req.StartMonth, req.EndMonth}

	// 如果有分类条件
	if req.CategoryId != 0 {
		query += " AND category_id = ?"
		args = append(args, req.CategoryId)
	}

	// 排序
	query += " ORDER BY amount DESC;"

	// 执行查询
	err := global.GVA_DB.Raw(query, args...).Find(&results).Error

	return results, err
}
