import { markRaw } from 'vue'
import { Money, TrendCharts, Wallet } from '@element-plus/icons-vue'

export const TASK_MIN_YEAR = 2025
export const MISER_MIN_YEAR = 2023
export const MISER_MIN_MONTH = 8

const TRANSACTION_TYPE_BALANCE = 0 // 余额
const TRANSACTION_TYPE_INCOME = 1 // 收入
const TRANSACTION_TYPE_EXPENSE = 2 //支出

const LOAN_FUND_STATUS_LOANING = 1 // 待还款
const LOAN_FUND_STATUS_REPAYING = 2 // 部分还
const LOAN_FUND_STATUS_REPAID = 3 // 已结清

export const miserCfgMap = {
  balance: {
    name: 'balance',
    label: '结余',
    transactionType: TRANSACTION_TYPE_BALANCE,
    color: '#00ccc6',
    colorTo: '#33d9d5',
    icon: markRaw(Wallet),
    tagType: 'primary'
  },
  income: {
    name: 'income',
    label: '收入',
    transactionType: TRANSACTION_TYPE_INCOME,
    color: '#67c23a',
    colorTo: '#85ce61',
    icon: markRaw(Money),
    tagType: 'success'
  },
  expense: {
    name: 'expense',
    label: '支出',
    transactionType: TRANSACTION_TYPE_EXPENSE,
    color: '#f56c6c',
    colorTo: '#f78989',
    icon: markRaw(TrendCharts),
    tagType: 'danger'
  }
}

export const miserTxnCfgMap = {
  [TRANSACTION_TYPE_BALANCE]: miserCfgMap.balance,
  [TRANSACTION_TYPE_INCOME]: miserCfgMap.income,
  [TRANSACTION_TYPE_EXPENSE]: miserCfgMap.expense
}

export const miserLoanCfgMap = {
  loaning: LOAN_FUND_STATUS_LOANING,
  repaying: LOAN_FUND_STATUS_REPAYING,
  repaid: LOAN_FUND_STATUS_REPAID
}
