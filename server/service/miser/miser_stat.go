package miser

import (
	"fmt"
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/miser/dto"
	"github.com/springbear2020/self-hub/server/model/miser/request"
	"golang.org/x/sync/errgroup"
	"math"
	"strconv"
	"strings"
)

type MiserStatService struct{}

const (
	TransactionTypeIncome  = 1
	TransactionTypeExpense = 2
)

var miserCategoryService = MiserCategoryService{}

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

func (s *MiserStatService) GetYearStat(uid uint, req *request.MiserStat) ([]dto.MiserStatYearAmount, error) {
	var results []dto.MiserStatYearAmount
	err := global.GVA_DB.Raw(`
	SELECT DATE_FORMAT(date, '%Y')                                                               AS year,
		   SUM(IF(transaction_type = 1, amount, 0))                                              AS income,
		   SUM(IF(transaction_type = 2, amount, 0))                                              AS expense,
		   (SUM(IF(transaction_type = 1, amount, 0)) - SUM(IF(transaction_type = 2, amount, 0))) AS balance
	FROM miser_transactions
	WHERE user_id = ?
	  AND DATE_FORMAT(date, '%Y') BETWEEN ? AND ?
	GROUP BY year
	ORDER BY year;
	`, uid, req.StartMonth[:4], req.EndMonth[:4],
	).Find(&results).Error
	return results, err
}

func (s *MiserStatService) GetMonthStat(uid uint, req *request.MiserStat) ([]dto.TableData, error) {
	var g errgroup.Group
	income := dto.TableData{Header: "收入明细"}
	expense := dto.TableData{Header: "支出明细"}

	g.Go(func() error {
		var err error
		income.Rows, income.Cols, err = s.getCategoryTableData(uid, req, TransactionTypeIncome)
		return err
	})

	g.Go(func() error {
		var err error
		expense.Rows, expense.Cols, err = s.getCategoryTableData(uid, req, TransactionTypeExpense)
		return err
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return []dto.TableData{income, expense}, nil
}

func (s *MiserStatService) GetLineStat(uid uint, req *request.MiserStat) ([]dto.MiserStatMonthAmount, error) {
	var results []dto.MiserStatMonthAmount
	err := global.GVA_DB.Raw(`
	SELECT DATE_FORMAT(date, '%Y-%m')                                                            AS month,
		   SUM(IF(transaction_type = 1, amount, 0))                                              AS income,
		   SUM(IF(transaction_type = 2, amount, 0))                                              AS expense,
		   (SUM(IF(transaction_type = 1, amount, 0)) - SUM(IF(transaction_type = 2, amount, 0))) AS balance
	FROM miser_transactions
	WHERE user_id = ?
	  AND DATE_FORMAT(date, '%Y-%m') BETWEEN ? AND ?
	GROUP BY month
	ORDER BY month;
	`, uid, req.StartMonth, req.EndMonth,
	).Find(&results).Error
	return results, err
}

func (s *MiserStatService) GetPieStat(uid uint, req *request.MiserStat) (map[string][]dto.MiserStatCategoryAmount, error) {
	var g errgroup.Group
	var incomes, expenses []dto.MiserStatCategoryAmount

	g.Go(func() error {
		var err error
		incomes, err = s.getCategoryAmount(uid, TransactionTypeIncome, req)
		return err
	})

	g.Go(func() error {
		var err error
		expenses, err = s.getCategoryAmount(uid, TransactionTypeExpense, req)
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
		totalIncome, err = s.getTotalAmount(uid, TransactionTypeIncome, req)
		return err
	})

	g.Go(func() error {
		var err error
		totalExpense, err = s.getTotalAmount(uid, TransactionTypeExpense, req)
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
