<script setup>
  import { ref, onMounted, onBeforeUnmount } from 'vue'
  import * as echarts from 'echarts'
  import {
    getCardStat,
    getCategoryStat,
    getLineStat,
    getPieStat
  } from '@/api/miser/miser_stat'
  import { miserCfgMap, miserTxnCfgMap } from '@/constants/miser'
  import { useMiserCategoryStore, useMiserStatStore } from '@/pinia'
  import { formatAmountCurrency } from '@/utils/format'
  import { storeToRefs } from 'pinia'

  defineOptions({ name: 'PieStat' })
  const emits = defineEmits(['open'])

  // 常量
  const {
    transactionType: expenseType,
    color: expenseColor,
    label: expenseLabel
  } = miserCfgMap.expense
  const {
    transactionType: balanceType,
    color: balanceColor,
    label: balanceLabel
  } = miserCfgMap.balance

  // 状态管理，用 storeToRefs 让解构后的值保持响应式
  const catStore = useMiserCategoryStore()
  const statStore = useMiserStatStore()
  const { fetchData, subscribe, unsubscribe } = statStore
  const { startMonth, endMonth } = storeToRefs(statStore)

  // 图表逻辑
  let incomeInstance, expenseInstance, balanceInstance
  const incomeRef = ref()
  const expenseRef = ref()
  const balanceRef = ref()

  const initChart = () => {
    if (!incomeRef.value || !expenseRef.value || !balanceRef.value) {
      return
    }

    if (incomeInstance || expenseInstance || balanceInstance) {
      disposeChart()
    }

    incomeInstance = echarts.init(incomeRef.value)
    expenseInstance = echarts.init(expenseRef.value)
    balanceInstance = echarts.init(balanceRef.value)

    incomeInstance.on('click', doPaymentClick)
    expenseInstance.on('click', doPaymentClick)
    balanceInstance.on('click', doBalanceClick)
  }

  const disposeChart = () => {
    if (!incomeInstance || !expenseInstance || !balanceInstance) {
      return
    }

    incomeInstance?.dispose()
    expenseInstance?.dispose()
    balanceInstance?.dispose()
    incomeInstance = null
    expenseInstance = null
    balanceInstance = null
  }

  const updateChart = (chartData = {}) => {
    if (!incomeInstance || !expenseInstance || !balanceInstance) {
      return
    }

    const { incomeData, expenseData, balanceData } = chartData

    const createPieOption = ({ title, data, color }) => ({
      title: {
        text: title,
        left: 'center',
        top: 10,
        textStyle: { fontSize: 16, fontWeight: 600 }
      },
      tooltip: {
        trigger: 'item',
        formatter: ({ name, value, percent }) =>
          `${name}: ${formatAmountCurrency(value)} (${percent}%)`
      },
      legend: { bottom: 0, type: 'scroll' },
      series: [
        {
          type: 'pie',
          radius: ['40%', '70%'],
          center: ['50%', '50%'],
          avoidLabelOverlap: false,
          itemStyle: { borderRadius: 6, borderColor: '#fff', borderWidth: 2 },
          label: { show: false, position: 'center' },
          emphasis: { label: { show: true, fontSize: 16, fontWeight: 'bold' } },
          labelLine: { show: false },
          color,
          data
        }
      ]
    })

    incomeInstance.clear()
    expenseInstance.clear()
    balanceInstance.clear()

    incomeInstance.setOption(createPieOption(incomeData))
    expenseInstance.setOption(createPieOption(expenseData))
    balanceInstance.setOption(createPieOption(balanceData))
  }

  // 收入占比 & 支出占比
  const doPaymentClick = async (pieData) => {
    const { categoryId, name, value, transactionType } = pieData.data
    const { code, data } = await fetchData(
      getCategoryStat.name,
      getCategoryStat,
      { categoryId }
    )
    if (code !== 0 || !Array.isArray(data)) {
      return
    }

    const { color, colorTo } = miserTxnCfgMap[transactionType]

    const chartData = {
      title: `${startMonth.value} 至 ${endMonth.value}『${name}』${formatAmountCurrency(value)}`,
      xData: data.map(({ month }) => month),
      yData: data.map(({ amount }) => amount),
      itemColor: color,
      areaColor: colorTo
    }

    emits('open', chartData)
  }

  // 支出/结余
  const doBalanceClick = async (pieData) => {
    const { code, data } = await fetchData(getLineStat.name, getLineStat)
    if (code !== 0 || !Array.isArray(data)) {
      return
    }

    const { transactionType, value, name } = pieData.data
    const { color, colorTo } = miserTxnCfgMap[transactionType]
    const key = miserTxnCfgMap[transactionType].name

    const chartData = {
      title: `${startMonth.value} 至 ${endMonth.value}『${name}』${formatAmountCurrency(value)}`,
      xData: data.map(({ month }) => month),
      yData: data.map((item) => item[key]),
      itemColor: color,
      areaColor: colorTo
    }

    emits('open', chartData)
  }

  // 数据加载
  const fetchAndRender = async () => {
    const pieResp = await fetchData(getPieStat.name, getPieStat)
    if (pieResp.code !== 0 || !pieResp.data) {
      return
    }
    const cardResp = await fetchData(getCardStat.name, getCardStat)
    if (cardResp.code !== 0 || !cardResp.data) {
      return
    }

    const { incomes, expenses } = pieResp.data
    const { expense, balance } = cardResp.data

    const mapCategoryData = (list) =>
      list.map(({ amount, category }) => ({
        name: catStore.dataMap[category],
        value: amount,
        categoryId: category,
        transactionType: catStore.catTxnMap[category]
      }))

    const incomeData = {
      title: '收入占比',
      data: mapCategoryData(incomes),
      color: []
    }

    const expenseData = {
      title: '支出占比',
      data: mapCategoryData(expenses),
      color: []
    }

    const balanceData = {
      title: '支出/结余',
      data: [
        { name: expenseLabel, value: expense, transactionType: expenseType },
        { name: balanceLabel, value: balance, transactionType: balanceType }
      ],
      color: [expenseColor, balanceColor]
    }

    updateChart({ incomeData, expenseData, balanceData })
  }

  // 生命周期
  const handleResize = () => {
    incomeInstance?.resize()
    expenseInstance?.resize()
    balanceInstance?.resize()
  }

  onMounted(async () => {
    initChart()
    await catStore.init()
    await fetchAndRender()
    subscribe(getPieStat.name, fetchAndRender)
    subscribe(getCardStat.name, fetchAndRender)
    window.addEventListener('resize', handleResize)
  })

  onBeforeUnmount(() => {
    window.removeEventListener('resize', handleResize)
    unsubscribe(getPieStat.name, fetchAndRender)
    unsubscribe(getCardStat.name, fetchAndRender)
    disposeChart()
  })
</script>

<template>
  <div class="chart-wrapper">
    <div ref="incomeRef" class="chart" />
    <div ref="expenseRef" class="chart" />
    <div ref="balanceRef" class="chart" />
  </div>
</template>

<style scoped lang="scss">
  .chart-wrapper {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
    margin-top: 8px;

    .chart {
      flex: 1;
      min-width: 300px;
      height: 320px;
      background: #fff;
      border-radius: 16px;
      box-shadow: 0 4px 14px rgba(0, 0, 0, 0.05);
      padding: 10px;
    }
  }
</style>
