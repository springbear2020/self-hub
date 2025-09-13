<script setup>
  import { ref, onMounted, onBeforeUnmount } from 'vue'
  import * as echarts from 'echarts'
  import { getLineStat, getMonthTransactionStat } from '@/api/miser/miser_stat'
  import { miserCfgMap, miserTxnCfgMap } from '@/constants/miser'
  import { useMiserCategoryStore, useMiserStatStore } from '@/pinia'
  import { formatAmountCurrency } from '@/utils/format'

  defineOptions({ name: 'LineStat' })
  const emits = defineEmits(['open'])

  // 常量
  const legends = [
    { ...miserCfgMap.income },
    { ...miserCfgMap.expense },
    { ...miserCfgMap.balance }
  ]

  // 状态管理
  const catStore = useMiserCategoryStore()
  const statStore = useMiserStatStore()
  const { fetchData, subscribe, unsubscribe } = statStore

  // 图表逻辑
  let chartInstance
  const chartRef = ref()

  const initChart = () => {
    if (!chartRef.value) {
      return
    }

    if (chartInstance) {
      disposeChart()
    }

    chartInstance = echarts.init(chartRef.value)

    chartInstance.on('click', openDialog)
  }

  const disposeChart = () => {
    if (!chartInstance) {
      return
    }

    chartInstance.dispose()
    chartInstance = null
  }

  const updateChart = (chartData = {}) => {
    if (!chartInstance) {
      return
    }

    const { names, xData, yData, colors } = chartData

    const series = names.map((name, index) => {
      return {
        name,
        type: 'line',
        smooth: true,
        data: yData[index],
        symbol: 'circle',
        symbolSize: 4,
        itemStyle: { color: colors[index] },
        lineStyle: { width: 2 }
      }
    })

    chartInstance.clear()
    chartInstance.setOption({
      tooltip: {
        trigger: 'axis',
        valueFormatter: (value) => formatAmountCurrency(value)
      },
      legend: {
        data: names,
        top: '4%'
      },
      grid: { left: '2%', right: '2%', bottom: '4%', containLabel: true },
      xAxis: {
        type: 'category',
        boundaryGap: false,
        data: xData,
        axisLabel: { rotate: 30 }
      },
      yAxis: {
        type: 'value',
        splitLine: { lineStyle: { type: 'dashed', color: '#ddd' } }
      },
      series
    })
  }

  // 对话框
  const openDialog = async ({ name: month, componentIndex, value: amount }) => {
    const { transactionType, label } = legends[componentIndex]
    if (!transactionType) {
      return
    }

    const params = {
      startMonth: month,
      endMonth: month,
      transactionType: transactionType
    }
    const { code, data } = await fetchData(
      getMonthTransactionStat.name,
      getMonthTransactionStat,
      params
    )
    if (code !== 0 || !Array.isArray(data)) {
      return
    }

    const filterSorted = data
      .filter((item) => item.amount > 0)
      .map((item) => ({
        ...item,
        amount: item.amount,
        categoryStr: catStore.dataMap[item.category]
      }))
      .sort((a, b) => b.amount - a.amount)

    const { color, colorTo } = miserTxnCfgMap[transactionType]

    const chartData = {
      title: `${month}『${label}』${formatAmountCurrency(amount)}`,
      xData: filterSorted.map(({ categoryStr }) => categoryStr),
      yData: filterSorted.map(({ amount }) => amount),
      itemColor: color,
      areaColor: colorTo
    }

    emits('open', chartData)
  }

  // 数据加载
  const fetchAndRender = async () => {
    const { code, data } = await fetchData(getLineStat.name, getLineStat)
    if (code !== 0 || !Array.isArray(data)) {
      return
    }

    const chartData = {
      names: legends.map(({ label }) => label),
      xData: data.map(({ month }) => month),
      yData: [
        data.map(({ income }) => income),
        data.map(({ expense }) => expense),
        data.map(({ balance }) => balance)
      ],
      colors: legends.map(({ color }) => color)
    }

    updateChart(chartData)
  }

  // 生命周期
  const handleResize = () => {
    chartInstance?.resize()
  }

  onMounted(async () => {
    initChart()
    await catStore.init()
    await fetchAndRender()
    subscribe(getLineStat.name, fetchAndRender)
    window.addEventListener('resize', handleResize)
  })

  onBeforeUnmount(() => {
    window.removeEventListener('resize', handleResize)
    unsubscribe(getLineStat.name, fetchAndRender)
    disposeChart()
  })
</script>

<template>
  <div class="chart-wrapper">
    <div ref="chartRef" class="chart" />
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
