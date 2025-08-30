<script setup>
  import { onBeforeUnmount, onMounted, ref, watch } from 'vue'
  import * as echarts from 'echarts'
  import { getLineStat } from '@/api/miser/miser_stat'
  import config from '@/core/config'

  const props = defineProps({
    startMonth: { type: String, required: true },
    endMonth: { type: String, required: true }
  })

  // 数据加载
  const dataList = ref([])
  const fetchAndRender = async () => {
    const params = {
      startMonth: props.startMonth,
      endMonth: props.endMonth
    }
    const { code, data } = await getLineStat(params)
    dataList.value = code === 0 ? data : []
    doRender()
  }

  // 图表渲染
  let chartInstance
  const chartRef = ref()
  const doRender = () => {
    chartInstance.clear()

    const months = dataList.value.map((i) => i.month)
    const incomes = dataList.value.map((i) => i.income)
    const expenses = dataList.value.map((i) => i.expense)
    const balances = dataList.value.map((i) => i.balance)

    chartInstance.setOption({
      tooltip: {
        trigger: 'axis'
      },
      legend: {
        data: ['收入', '支出', '结余'],
        top: '4%'
      },
      grid: {
        left: '2%',
        right: '2%',
        bottom: '4%',
        containLabel: true
      },
      xAxis: {
        type: 'category',
        boundaryGap: false,
        data: months
      },
      yAxis: {
        type: 'value'
      },
      series: [
        {
          name: '收入',
          type: 'line',
          smooth: true,
          data: incomes,
          color: config.miser.income.color
        },
        {
          name: '支出',
          type: 'line',
          smooth: true,
          data: expenses,
          color: config.miser.expense.color
        },
        {
          name: '结余',
          type: 'line',
          smooth: true,
          data: balances,
          color: config.miser.balance.color
        }
      ]
    })
  }

  onMounted(() => {
    chartInstance = echarts.init(chartRef.value)

    fetchAndRender()
  })

  onBeforeUnmount(() => {
    chartInstance?.dispose()
  })

  watch(
    () => [props.startMonth, props.endMonth],
    () => {
      fetchAndRender()
    }
  )
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
